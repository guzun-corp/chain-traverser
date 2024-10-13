package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"

	"chain-traverser/internal/blockchain/eth"
	"chain-traverser/internal/config"
	"chain-traverser/internal/storage/redis"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

const WEI_IN_ETH = 1000000000000000000

type Indexer struct {
	client *eth.EthClient
	redis  *redis.RedisClient
	cfg    *config.Config
}

func NewIndexer(cfg *config.Config) (*Indexer, error) {
	redis := redis.NewClient(&cfg.Redis)
	client, err := eth.NewEthClient(&cfg.Eth)
	if err != nil {
		return nil, fmt.Errorf("error connecting to Ethereum node: %w", err)
	}
	return &Indexer{
		client: client,
		redis:  redis,
		cfg:    cfg,
	}, nil
}

// if address has more then 10k transactions and no labels, ask to enrich
func (i *Indexer) askToEnrichAddress(address string) {
	cnt, _ := i.redis.GetAddressTxNumber(&address)
	if cnt < 10_000 {
		return
	}
	labels, _ := i.redis.GetAddressLabels(&address)
	if labels != nil {
		return
	}
	i.redis.SendAddress(address)
}

func (i *Indexer) handleBlock(blockNumber *big.Int, ctx context.Context) *map[string]int64 {
	block, err := i.client.Client.BlockByNumber(ctx, blockNumber)
	if err != nil {
		log.Err(err).Msgf("fetch block %d by number error", blockNumber)
		// time.Sleep(5 * time.Second)
		// resultChan <- nil
		return nil
	}
	transMap := make(map[string]int64)
	blockTime := block.Time()
	blockEthPriceUsd, err := eth.GetTokenPrice(blockTime, i.redis, eth.ETH)
	if err != nil || blockEthPriceUsd == nil {
		log.Err(err).Msgf("error getting price for block: %d", blockNumber)
		return nil
	}
	blob := ""
	for _, tx := range block.Transactions() {
		if tx.To() == nil {
			log.Warn().Msgf("to is nil, skip; tx: %s", tx.Hash().Hex())
			continue
		}
		if from, err := types.Sender(types.NewLondonSigner(big.NewInt(1)), tx); err == nil {
			toHash := tx.To().Hex()
			_, toExists := transMap[toHash]
			if !toExists {
				transMap[toHash] = 1
			} else {
				transMap[toHash] += 1
			}

			fromHash := from.Hex()
			_, fromExists := transMap[fromHash]
			if !fromExists {
				transMap[fromHash] = 1
			} else {
				transMap[fromHash] += 1
			}
			value := tx.Value()
			var usdOnDay decimal.Decimal
			if value != nil && value.Cmp(big.NewInt(0)) == 1 {
				eth := decimal.NewFromBigInt(value, 0).Div(decimal.NewFromInt(WEI_IN_ETH))
				usd := eth.Mul(*blockEthPriceUsd)
				usdOnDay = usd.RoundBank(2)
				// set erc20 values to 0
				blob += fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s;%s\n", fromHash, tx.Hash().Hex(), toHash, value, usdOnDay.String(), "nil", "0", "0")
				i.askToEnrichAddress(fromHash)
				i.askToEnrichAddress(toHash)
			} else {
				// if valie is 0, check if it's erc20
				erc20tx, err := i.client.HandleERC20(*tx)
				if err == nil && erc20tx != nil {
					tokenPrice, err := eth.GetTokenPrice(blockTime, i.redis, erc20tx.Ticker)
					if err != nil || tokenPrice == nil {
						continue
					}
					usd := erc20tx.Value.Mul(*tokenPrice)
					usdOnDay = usd.RoundBank(2)

					blob += fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s;%s\n", fromHash, tx.Hash().Hex(), erc20tx.To, "0", "0", erc20tx.Ticker, erc20tx.Value, usdOnDay)
					i.askToEnrichAddress(fromHash)
					i.askToEnrichAddress(toHash)
				}

			}

		}
	}
	go i.redis.AddBlock(blockNumber, &blob)
	return &transMap
}

func (i *Indexer) getNextBlockNumber() (*big.Int, error) {
	blockNumber, err := i.redis.GetLastBlockNumber()
	if err != nil {
		return nil, err
	}

	if err != nil && err.Error() != "redis: nil" {
		return nil, err
	}
	var blockNumberToHandle *big.Int
	if blockNumber == nil {
		blockNumberToHandle = big.NewInt(i.cfg.Indexer.StartBlockNumber)
	} else {
		blockNumberToHandle = big.NewInt(*blockNumber + 1)
	}
	return blockNumberToHandle, nil

}

func (i *Indexer) processBlocks(ctx context.Context) error {
	prev_time := time.Now()
	blockCount := 0

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			blockNumber, err := i.getNextBlockNumber()
			if err != nil {
				log.Err(err).Msg("error getting next block number")
				time.Sleep(5 * time.Second)
				continue
			}

			result := i.handleBlock(blockNumber, ctx)
			if result == nil {
				log.Info().Msgf("result is nil, skip bulk | block: %s", blockNumber)
				time.Sleep(17 * time.Second)
				continue
			}

			go i.redis.UpdateCounters(*result)
			go i.redis.AppendBlockNumbers(*result, blockNumber)

			err = i.redis.UpdateLastBlockNumber(blockNumber)
			if err != nil {
				log.Err(err).Msg("error during updating next block number")
				time.Sleep(5 * time.Second)
				continue
			}

			blockCount++
			if blockCount%100 == 0 {
				i.logProgress(blockNumber, prev_time)
				prev_time = time.Now()
			}
		}
	}
}

func (i *Indexer) logProgress(blockNumber *big.Int, prevTime time.Time) {
	blocksPerSec := 100 / (time.Since(prevTime).Seconds())
	log.Info().Msgf("%s | Blocks %d | blocks p/s: %f", time.Now().UTC(), blockNumber, blocksPerSec)
}

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	cfg, err := config.NewConfig()

	if err != nil {
		log.Err(err).Msg("error reading config")
		return
	}

	indexer, err := NewIndexer(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create indexer")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		<-stop
		log.Info().Msg("Interrupt received, terminating...")
		cancel()
	}()

	if err := indexer.processBlocks(ctx); err != nil && err != context.Canceled {
		log.Fatal().Err(err).Msg("Error processing blocks")
	}
}
