package traverser

import (
	"chain-traverser/internal/storage/redis"
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

const TRAVERSE_MAX_DEGREE = 300

type CntRes struct {
	key *string
	cnt int64
}

type CntErr struct {
	key *string
	err error
}

func setCounters(addrs *map[string]Addr, storage *redis.RedisClient) error {
	resChan := make(chan CntRes, len(*addrs))
	errChan := make(chan CntErr, len(*addrs))

	for key, _ := range *addrs {
		go func(key *string) {
			cnt, err := storage.GetAddressTxNumber(key)
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
			v.NeedTraverse = false
			(*addrs)[*errRes.key] = v
			log.Err(errRes.err).Msgf("Cant set counter for %s", key)
		case res := <-resChan:
			v := (*addrs)[*res.key]
			v.Cnt = res.cnt
			if v.Cnt > TRAVERSE_MAX_DEGREE {
				v.NeedTraverse = false
			}
			(*addrs)[*res.key] = v
		}
	}

	return nil
}

func getBlocks(addrs *map[string]Addr, storage *redis.RedisClient) (*[]string, error) {
	traverseAddrs := []string{}
	for key, addr := range *addrs {
		if addr.NeedTraverse {
			traverseAddrs = append(traverseAddrs, key)
		}
	}

	blocksSet := make(map[string]bool)

	resChan := make(chan *[]string, len(*addrs))
	errChan := make(chan error, len(*addrs))
	for _, key := range traverseAddrs {
		go func(key *string) {
			bSlice, err := storage.GetAddressBlocks(key)
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
			log.Err(err).Msgf("Cant set block list for %d", key)
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

func getTransactionsByBlockNumber(blockNumber string, addrs *map[string]Addr, storage *redis.RedisClient, fromBlock int, toBlock int, limiter *AtomicLimiter) (*[]Tx, error) {
	blockNumberInt, err := strconv.Atoi(blockNumber)
	if err != nil {
		log.Err(err).Msgf("Failed to convert blockNumber to int: %s", blockNumber)
		return nil, err
	}
	txs := []Tx{}

	if fromBlock > blockNumberInt || toBlock < blockNumberInt {
		return &txs, nil
	}

	block, err := storage.GetBlock(&blockNumber)
	if err != nil {
		log.Err(err).Msgf("getTransactionsByBlockNumber can't get block: %s", blockNumber)
		return nil, err
	}
	for _, tx := range strings.Split(*block, "\n") {
		if limiter.IsExceed() {
			log.Warn().Msgf("limiter is exceed")
			break
		}
		if tx == "" {
			continue
		}
		vals := strings.Split(tx, ";")
		from := vals[0]
		to := vals[2]
		txHash := vals[1]

		fromAddr, existsFrom := (*addrs)[from]
		toAddr, existsTo := (*addrs)[to]
		// the following conditions exclude transactions that are not interesting for us
		// we skip traversing over dex explicitly.
		// hovewer, if one of the addresses had interaction with dex we got the dex address.
		// We prefer to leave dex address with interaction transaction in graph without other dexs transactions.
		// skip them
		if !existsFrom && !existsTo {
			continue
		}
		if existsFrom && existsTo && !fromAddr.NeedTraverse && !toAddr.NeedTraverse {
			continue
		}
		if existsFrom && !fromAddr.NeedTraverse && !existsTo {
			continue
		}
		if existsTo && !toAddr.NeedTraverse && !existsFrom {
			continue
		}

		ethAmount, _ := decimal.NewFromString(vals[3])
		ethAmount = ethAmount.Div(decimal.NewFromInt(1e18))
		ethAmountUsdOnDay, _ := decimal.NewFromString(vals[4])

		totalUsdFlow := ethAmountUsdOnDay
		flowByCurrency := make(map[string]decimal.Decimal)
		flowByCurrency["ETH"] = ethAmount
		erc20 := vals[5]
		if erc20 != "nil" {
			erc20AmountUsdOnDay, _ := decimal.NewFromString(vals[7])
			erc20Amount, _ := decimal.NewFromString(vals[6])
			totalUsdFlow = ethAmountUsdOnDay.Add(erc20AmountUsdOnDay)
			flowByCurrency[erc20] = erc20Amount
		}

		txs = append(txs, Tx{
			From:           from,
			To:             to,
			TxHash:         txHash,
			TotalUsdFlow:   totalUsdFlow,
			FlowByCurrency: flowByCurrency,
		})
		limiter.Consume()
	}

	return &txs, nil
}

func getAddressTransactions(addrs map[string]Addr, redis *redis.RedisClient, ctx context.Context, depth int, fromBlock int, toBlock int, limiter *AtomicLimiter) (*[]Tx, error) {
	log.Debug().Msgf("getAddressTransactions %d", depth)
	if depth == 0 {
		return nil, nil
	}

	err := setCounters(&addrs, redis)
	if err != nil {
		return nil, err
	}
	log.Debug().Msgf("addr count: %d", len(addrs))

	blocks, err := getBlocks(&addrs, redis)
	if err != nil {
		log.Err(err).Msgf("getAddressTransactions failed on getting blocks | address: %s", addrs)
		return nil, err
	}
	log.Info().Msgf("depth: %d", depth)
	log.Info().Msgf("block count: %d", len(*blocks))

	var txs []Tx
	resultChan := make(chan []Tx, len(*blocks))
	errChan := make(chan error, len(*blocks))

	for _, block := range *blocks {
		go func(block string) {
			local_txs, err := getTransactionsByBlockNumber(block, &addrs, redis, fromBlock, toBlock, limiter)
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
			log.Err(err).Msgf("getTransactionsByBlockNumber | Error: %s", err)
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
			nextAddrs[tx.From] = Addr{Hash: tx.From, Cnt: 0, NeedTraverse: true}
		}
		_, exists = addrs[tx.To]
		if !exists {
			nextAddrs[tx.To] = Addr{Hash: tx.To, Cnt: 0, NeedTraverse: true}
		}
	}
	txs2, err := getAddressTransactions(nextAddrs, redis, ctx, depth-1, fromBlock, toBlock, limiter)
	if err != nil {
		return nil, err
	}
	txs = append(txs, *txs2...)
	return &txs, nil
}

func CollectBFS(address string, depth int, fromBlock int, toBlock int, redis *redis.RedisClient) (*Graph, error) {
	log.Info().Msgf("CollectBFS: %s, %d", address, depth)

	ctx := context.Background()

	start := time.Now()

	addrs := make(map[string]Addr)
	addrs[address] = Addr{Hash: address, Cnt: -1, NeedTraverse: true}
	limiter := NewLimiter()
	txs, err := getAddressTransactions(addrs, redis, ctx, depth, fromBlock, toBlock, limiter)

	log.Info().Msgf("got all transactions in %s", time.Since(start))

	if err != nil {
		log.Err(err).Msgf("getAddressTransactions failed %s", address)
		return nil, err
	}
	log.Info().Msgf("got %d transactions", len(*txs))

	uTrsx := make(map[string]Tx)
	for _, tx := range *txs {
		uTrsx[tx.TxHash] = tx
	}
	log.Info().Msgf("got %d unique transactions", len(uTrsx))
	return &Graph{Addrs: &addrs, Txs: &uTrsx}, nil
}
