package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const TRAVERSE_MAX_DEGREE = 70

type Tx struct {
	From   string
	To     string
	TxHash string
}

type Addr struct {
	hash         string
	cnt          int64
	needTraverse bool
}

func getAddressTxNumber(addr *string, redis *redis.Client, ctx context.Context) (int64, error) {
	key := fmt.Sprintf("c:%s", *addr)
	val, err := redis.Get(ctx, key).Int64()
	if err != nil {
		log.Err(err).Msg("Cant get counter")
	}
	log.Debug().Msgf("got counter %d for %s", val, *addr)
	return val, nil
}

type CntRes struct {
	key *string
	cnt int64
}
type CntErr struct {
	key *string
	err error
}

func setCounters(addrs *map[string]Addr, redis *redis.Client, ctx context.Context) error {
	resChan := make(chan CntRes, len(*addrs))
	errChan := make(chan CntErr, len(*addrs))
	// log.Debug().Msgf("addrs %s", addrs)
	for key, _ := range *addrs {
		go func(key *string) {
			cnt, err := getAddressTxNumber(key, redis, ctx)
			if err != nil {
				errChan <- CntErr{key: key, err: err}
			} else {
				resChan <- CntRes{key: key, cnt: cnt}
			}
		}(&key)
	}

	for key, _ := range *addrs {
		select {
		case errRes := <-errChan:
			v := (*addrs)[*errRes.key]
			v.needTraverse = false
			(*addrs)[*errRes.key] = v
			log.Err(errRes.err).Msgf("Cant set counter for %s", key)
		case res := <-resChan:
			v := (*addrs)[*res.key]
			v.cnt = res.cnt
			if v.cnt > TRAVERSE_MAX_DEGREE {
				v.needTraverse = false
				// log.Debug().Msgf("skip key %s, addr %s, cnt %d", *res.key, v.hash, v.cnt)
			}
			(*addrs)[*res.key] = v
		}
	}

	return nil
}

func getAddressBlocks(addr *string, redis *redis.Client, ctx context.Context) (*[]string, error) {
	key := fmt.Sprintf("b:%s", *addr)
	log.Info().Msgf("getting blocks for %s", key)
	slice, err := redis.LRange(ctx, key, 0, 1000).Result()
	if err != nil {
		log.Err(err).Msg("Cant get block list")
		return nil, err
	}
	// log.Debug().Msgf("blocks %s", slice)
	log.Debug().Msgf("blocks %d", len(slice))
	return &slice, nil
}

func getBlocks(addrs *map[string]Addr, redis *redis.Client, ctx context.Context) (*[]string, error) {
	// blocks, err := getAddressBlocks(addrs, redis, ctx)
	// if err != nil {
	// 	return nil, err
	// }
	traverseAddrs := []string{}
	for key, addr := range *addrs {
		if addr.needTraverse {
			traverseAddrs = append(traverseAddrs, key)
		}
	}

	blocksSet := make(map[string]bool)

	resChan := make(chan *[]string, len(*addrs))
	errChan := make(chan error, len(*addrs))
	for _, key := range traverseAddrs {
		go func(key *string) {
			bSlice, err := getAddressBlocks(key, redis, ctx)
			if err != nil {
				errChan <- err
			} else {
				resChan <- bSlice
			}
		}(&key)
	}
	for key, _ := range traverseAddrs {
		select {
		case err := <-errChan:
			log.Err(err).Msgf("Cant set block list for %s", key)
		case res := <-resChan:
			for _, block := range *res {
				blocksSet[block] = true
			}
		}
	}

	blocks := []string{}
	for key, _ := range blocksSet {
		blocks = append(blocks, key)
	}
	return &blocks, nil
}

func getTransactionsByBlockNumber(blockNumber string, address string, client *ethclient.Client, ctx context.Context) (*[]Tx, error) {
	bigInt := new(big.Int)
	bigInt.SetString(blockNumber, 10)

	block, err := client.BlockByNumber(ctx, bigInt)
	if err != nil {
		return nil, err
	}
	txs := []Tx{}
	// log.Debug().Msgf("block %s", blockNumber)

	for _, tx := range block.Transactions() {
		if tx.To() == nil {
			continue
		}
		if from, err := types.Sender(types.NewLondonSigner(big.NewInt(1)), tx); err == nil {
			if from.Hex() != address && tx.To().Hex() != address {
				continue
			}
			txs = append(txs, Tx{
				From:   from.Hex(),
				To:     tx.To().Hex(),
				TxHash: tx.Hash().Hex(),
			})
		}
	}
	// appendBlockTimestamp(transMap, blockNumber, redis, ctx)
	return &txs, nil
}

func getBlockRedis(blockNumber string, redis *redis.Client, ctx context.Context) (*string, error) {
	// start := time.Now()
	key := fmt.Sprintf("tx:%s", blockNumber)
	val, err := redis.Get(ctx, key).Result()
	if err != nil {
		log.Err(err).Msg("Cant get block")
		return nil, err
	}
	// log.Debug().Msgf("get block data %s in %s", blockNumber, time.Since(start))
	return &val, nil
}
func getTransactionsByBlockNumber2(blockNumber string, addrs *map[string]Addr, redis *redis.Client, ctx context.Context) (*[]Tx, error) {
	// start := time.Now()

	bigInt := new(big.Int)
	bigInt.SetString(blockNumber, 10)

	block, err := getBlockRedis(blockNumber, redis, ctx)
	if err != nil {
		return nil, err
	}
	txs := []Tx{}
	// log.Debug().Msgf("block %s", blockNumber)

	// for key, addr := range *addrs {
	// 	log.Debug().Msgf("addr %s %d %t", key, addr.cnt, addr.needTraverse)
	// }
	for _, tx := range strings.Split(*block, "\n") {
		if tx == "" {
			continue
		}
		vals := strings.Split(tx, ";")
		from := vals[0]
		to := vals[2]
		txHash := vals[1]

		// log.Debug().Msgf("tx %s %s %s", from, to, txHash)

		_, existsFrom := (*addrs)[from]
		_, existsTo := (*addrs)[to]
		if !existsFrom && !existsTo {
			continue
		}
		// log.Debug().Msgf("tx %s %s %s", from, to, txHash)
		txs = append(txs, Tx{
			From:   from,
			To:     to,
			TxHash: txHash,
		})
	}
	// appendBlockTimestamp(transMap, blockNumber, redis, ctx)
	// log.Debug().Msgf("get transactions by block %s in %s", blockNumber, time.Since(start))
	return &txs, nil
}

func getAddressTransactions(addrs map[string]Addr, redis *redis.Client, client *ethclient.Client, ctx context.Context, depth uint8) (*[]Tx, error) {
	log.Debug().Msgf("getAddressTransactions %d", depth)
	if depth == 0 {
		return nil, nil
	}

	err := setCounters(&addrs, redis, ctx)
	if err != nil {
		return nil, err
	}
	log.Debug().Msgf("addr count: %d", len(addrs))
	// for key, addr := range addrs {
	// 	log.Debug().Msgf("addr %s %d %t", key, addr.cnt, addr.needTraverse)
	// }
	blocks, err := getBlocks(&addrs, redis, ctx)
	if err != nil {
		return nil, err
	}
	log.Debug().Msgf("block count: %s", len(*blocks))
	// for block := range blocks {
	// 	log.Debug().Msgf("block %d", block)
	// }

	var txs []Tx
	resultChan := make(chan []Tx, len(*blocks))
	errChan := make(chan error, len(*blocks))

	for _, block := range *blocks {
		go func(block string) {
			// local_txs, err := getTransactionsByBlockNumber(block, address, client, ctx)
			local_txs, err := getTransactionsByBlockNumber2(block, &addrs, redis, ctx)
			if err != nil {
				errChan <- err
			} else {
				resultChan <- *local_txs
			}
		}(block)
	}

	for range *blocks {
		select {
		case err := <-errChan:
			return nil, err
		case result := <-resultChan:
			txs = append(txs, result...)
		}
	}
	if depth == 1 {
		return &txs, nil
	}

	// collect next bunch of addresses
	// skip addresses from previous iteration
	nextAddrs := make(map[string]Addr)
	for _, tx := range txs {
		_, exists := addrs[tx.From]
		if !exists {
			nextAddrs[tx.From] = Addr{hash: tx.From, cnt: 0, needTraverse: true}
		}
		_, exists = addrs[tx.To]
		if !exists {
			nextAddrs[tx.To] = Addr{hash: tx.To, cnt: 0, needTraverse: true}
		}
	}
	txs2, err := getAddressTransactions(nextAddrs, redis, client, ctx, depth-1)
	if err != nil {
		return nil, err
	}
	txs = append(txs, *txs2...)
	return &txs, nil
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	address := "0x5C46324785eF1B359a5546736bB84FAdc14ac550"

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Err(err)
		return
	}
	log.Info().Msg("Connected to Ethereum node")

	start := time.Now()

	addrs := make(map[string]Addr)
	addrs[address] = Addr{hash: address, cnt: -1, needTraverse: true}
	txs, err := getAddressTransactions(addrs, redis, client, ctx, 3)

	log.Debug().Msgf("got all transactions in %s", time.Since(start))

	if err != nil {
		log.Err(err).Msg("getAddressTransactions failed")
		return
	}
	log.Info().Msgf("got %d transactions", len(*txs))

	uTrsx := make(map[string]Tx)
	for _, tx := range *txs {
		uTrsx[tx.TxHash] = tx
	}
	log.Info().Msgf("got %d unique transactions", len(uTrsx))
	// for _, tx := range *txs {
	// 	fmt.Println(tx)
	// }
}
