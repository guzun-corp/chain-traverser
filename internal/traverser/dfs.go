package traverser

import (
	"chain-traverser/internal/config"
	"chain-traverser/internal/storage/redis"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

func addressBlocks(addr string, fromBlock int, toBlock int, redis *redis.RedisClient) (*[]int, error) {
	bSlice, err := redis.GetAddressBlocks(&addr)
	if err != nil {
		log.Err(err).Msgf("Cant get blocks for %s", addr)
		return nil, err
	}
	var filteredBlocks []int
	for _, block := range *bSlice {
		blockNumber, err := strconv.Atoi(block)
		if err != nil {
			log.Err(err).Msgf("Error converting block number to int: %d", blockNumber)
			continue
		}
		if blockNumber >= fromBlock && blockNumber <= toBlock {
			filteredBlocks = append(filteredBlocks, blockNumber)
		}
	}

	return &filteredBlocks, nil
}

func blockTransactions(blockNumber string, addr string, flow string, redis *redis.RedisClient) (*[]Tx, error) {
	bigInt := new(big.Int)
	bigInt.SetString(blockNumber, 10)

	block, err := redis.GetBlock(&blockNumber)
	if err != nil {
		log.Err(err).Msgf("blockTransactions can't get block: %s", blockNumber)
		return nil, err
	}
	txs := []Tx{}
	for _, tx := range strings.Split(*block, "\n") {
		if tx == "" {
			continue
		}
		vals := strings.Split(tx, ";")
		from := vals[0]
		to := vals[2]
		txHash := vals[1]

		if (flow == "input" || flow == "all") && to == addr ||
			(flow == "output" || flow == "all") && from == addr ||
			(flow == "all") && (from == addr || to == addr) {

			ethAmount, _ := decimal.NewFromString(vals[3])
			ethAmount = ethAmount.Div(decimal.NewFromInt(1e18))
			ethAmountUsdOnDay, _ := decimal.NewFromString(vals[4])
			erc20 := vals[5]

			totalUsdFlow := ethAmountUsdOnDay
			flowByCurrency := make(map[string]decimal.Decimal)
			flowByCurrency["ETH"] = ethAmount
			if erc20 != "nil" {
				erc20AmountUsdOnDay, _ := decimal.NewFromString(vals[7])
				erc20Amount, _ := decimal.NewFromString(vals[6])
				totalUsdFlow = ethAmountUsdOnDay.Add(erc20AmountUsdOnDay)
				flowByCurrency[erc20] = erc20Amount
			}
			// log.Debug().Msgf("tx %s %s %s", from, to, txHash)
			txs = append(txs, Tx{
				From:           from,
				To:             to,
				TxHash:         txHash,
				TotalUsdFlow:   totalUsdFlow,
				FlowByCurrency: flowByCurrency,
			})
		}

	}
	return &txs, nil
}

func getAddress(addr AddrWithDepth, redis *redis.RedisClient) (*Addr, error) {
	addrCnt, _ := redis.GetAddressTxNumber(&addr.hash)
	needTraverse := true
	if addr.depth != 0 && addrCnt > TRAVERSE_MAX_DEGREE {
		log.Debug().Msgf("skip address cause of degree = ", addrCnt)
		needTraverse = false
	}
	return &Addr{Hash: addr.hash, Cnt: addrCnt, NeedTraverse: needTraverse}, nil
}

func getTrxFrom(addr string, fromBlock int, toBlock int, flow string, redis *redis.RedisClient) (*[]Tx, error) {
	blocks, err := addressBlocks(addr, fromBlock, toBlock, redis)
	if err != nil {
		return nil, err
	}
	resultChan := make(chan []Tx, len(*blocks))
	errChan := make(chan error, len(*blocks))

	for _, block := range *blocks {
		go func(block int) {
			local_txs, err := blockTransactions(strconv.Itoa(block), addr, flow, redis)
			if err != nil {
				errChan <- err
			} else {
				resultChan <- *local_txs
			}
		}(block)
	}

	var txs []Tx

	for range *blocks {
		select {
		case err := <-errChan:
			log.Err(err).Msgf("getTransactionsByBlockNumber | Error: %s", err)
		case result := <-resultChan:
			txs = append(txs, result...)
		}
	}

	return &txs, nil
}

type AddrWithDepth struct {
	hash  string
	depth int
}

type ParamsDFS struct {
	Address        string
	Depth          int
	FromBlock      int
	ToBlock        int
	Flow           string
	GraphSizeLimit int
}

func (p ParamsDFS) String() string {
	return fmt.Sprintf("address: %s, depth: %d, fromBlock: %d, toBlock: %d, flow: %s, graphSizeLimit: %d",
		p.Address, p.Depth, p.FromBlock, p.ToBlock, p.Flow, p.GraphSizeLimit)
}

func CollectDFS(params ParamsDFS, redis *redis.RedisClient) (*Graph, error) {
	log.Info().Msgf("CollectDFS: %s", fmt.Sprintf("%+v", params))
	graph := &Graph{
		Addrs: &map[string]Addr{},
		Txs:   &map[string]Tx{},
	}
	firstAddr := AddrWithDepth{hash: params.Address, depth: 0}
	stack := []AddrWithDepth{firstAddr}

	depthCnt := 0
	// DFS logic
	addressCnter := 0
	for len(stack) > 0 && len(*graph.Addrs) < params.GraphSizeLimit && len(*graph.Txs) < params.GraphSizeLimit && depthCnt < params.Depth {
		// Pop the top address from the stack
		addr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if addr.depth >= params.Depth {
			continue
		}
		if _, exists := (*graph.Addrs)[addr.hash]; exists {
			continue
		}

		addrObj, err := getAddress(addr, redis)
		if err != nil {
			log.Err(err).Msgf("Cant get address %s", addr.hash)
			continue
		}
		(*graph.Addrs)[addr.hash] = *addrObj

		if !addrObj.NeedTraverse {
			continue
		}

		trxs, err := getTrxFrom(addr.hash, params.FromBlock, params.ToBlock, params.Flow, redis)
		if err != nil {
			log.Err(err).Msgf("Cant get transactions for %s", addr.hash)
			continue
		}

		if addr.hash == params.Address {
			addressCnter += 1
		}

		// Visit all transactions from this address
		for _, tx := range *trxs {
			(*graph.Txs)[tx.TxHash] = tx
			var addAddr string
			if tx.To == addr.hash {
				addAddr = tx.From
			} else {
				addAddr = tx.To
			}
			if addAddr == addr.hash {
				// handle case when address sends to itself
				continue
			}
			stack = append(stack, AddrWithDepth{hash: addAddr, depth: addr.depth + 1})
		}
		if len(stack)%100 == 0 {
			log.Info().Msgf("Stack length: %d", len(stack))
		}
		(*addrObj).NeedTraverse = false
	}
	if addressCnter > 2 {
		log.Warn().Msgf("address cnter = %d", addressCnter)
	}
	return graph, nil
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	cfg, err := config.NewConfig()

	if err != nil {
		log.Err(err).Msg("error reading config")
		return
	}
	redis := redis.NewClient(&cfg.Redis)

	// Example usage
	params := ParamsDFS{"startAddress", 3, 0, 20_000_000, "all", 5000}
	graph, err := CollectDFS(params, redis)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the collected graph
	log.Debug().Msgf("Collected Graph: %s", graph)
}
