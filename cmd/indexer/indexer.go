package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"

	"chain-traverser/internal/storage/redis"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const LAST_BLOCK_FILE = "last_handled_block_trans_counter.txt"

func getNextBlockNumber() (*big.Int, error) {
	f, err := os.Open(LAST_BLOCK_FILE)
	blockNumber := 0
	if err != nil {
		if os.IsNotExist(err) {
			blockNumber = 19050000
		} else {
			// Other error occurred while opening the file
			return nil, err
		}
	}
	defer f.Close()

	if blockNumber == 0 {
		_, err = fmt.Fscanf(f, "%d", &blockNumber)
		if err != nil {
			return nil, err
		}
	}
	return big.NewInt(int64(blockNumber + 1)), nil
}

func setBlockNumber(blockNumber *big.Int) error {
	f, err := os.Create(LAST_BLOCK_FILE)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, "%d", blockNumber)
	if err != nil {
		return err
	}

	return nil
}

func handleBlock(
	blockNumber *big.Int,
	client *ethclient.Client,
	redis *redis.RedisClient,
	ctx context.Context,
) *map[string]int64 {
	block, err := client.BlockByNumber(ctx, blockNumber)
	if err != nil {
		fmt.Printf("fetch block by number error: %s\n", err)
		// time.Sleep(5 * time.Second)
		// resultChan <- nil
		return nil
	}
	// fmt.Println("got block: %s\n", block.Time())
	transMap := make(map[string]int64)
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
			blob += fmt.Sprintf("%s;%s;%s\n", fromHash, tx.Hash().Hex(), toHash)
		}
	}
	go redis.AddBlock(blockNumber, &blob)
	return &transMap
}

// func appendBlockTimestamp(transMap map[string]int64, blockNumber *big.Int, redis *redis.Client, ctx context.Context) {
// 	redis.Set(ctx, timestamp, blockNumber)
// }

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	redis := redis.NewClient()

	ctx := context.Background()

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Err(err).Msg("error connecting to Ethereum node")
		return
	}
	log.Info().Msg("Connected to Ethereum node")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	prev_time := time.Now()

	i := 0
	for {
		select {
		case <-stop:
			// Handle interrupt, then return to terminate the goroutine
			log.Info().Msg("Interrupt received, terminating...")
			return
		default:
			// start := time.Now()

			blockNumber, err := getNextBlockNumber()
			if err != nil {
				log.Err(err).Msg("error getting next block number")
				return
			}

			if i%100 == 0 {

				bloksPerSec := 100 / (time.Since(prev_time).Seconds())

				log.Info().Msgf("%s | Blocks %d | blocks p/s: %f", time.Now().UTC(), blockNumber, bloksPerSec)

				prev_time = time.Now()
			}

			result := handleBlock(blockNumber, client, redis, ctx)
			if result == nil {
				log.Info().Msg("result is nil, skip bulk")
				time.Sleep(17 * time.Second)
				continue
			}
			// fmt.Printf("start storing block: %s\n", blockNumber)

			go redis.UpdateCounters(*result)
			go redis.AppendBlockNumbers(*result, blockNumber)

			// log.Debug().Msgf("handled block: %s\n", blockNumber)

			setBlockNumber(blockNumber)

			// elapse := time.Since(start)
			// fmt.Printf("last handled block: %s | b/s %f\n", lastHadled, float64(1)/elapse.Seconds())
			i += 1
			// time.Sleep(3 * time.Second)

		}
	}
}
