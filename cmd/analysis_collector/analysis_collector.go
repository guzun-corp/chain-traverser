package main

import (
	"context"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"

	"chain-traverser/internal/blockchain/eth"
	"chain-traverser/internal/storage/redis"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

const START_BLOCK_NUMBER = 19050000

const FINISH_BLOCK_NUMBER = 19131235

// const FINISH_BLOCK_NUMBER = 19050010

const WEI_IN_ETH = 1000000000000000000

// we have some failed transactions
// also we can't afford to get receipt for each transaction to identify failed ones
// the simple way to cut every transaction with value bigger then usdt capitalization
var MAX_USDT_LIMIT = decimal.NewFromInt(107049317263) // capitalization of USDT

type Wallet struct {
	Address               string
	TxTotal               int64
	TxIn                  int64
	TxOut                 int64
	TxTotalVolumeUsdOnDay float64
	TxInVolumeUsdOnDay    float64
	TxOutVolumeUsdOnDay   float64
	BalanceUsd            float64
	Currencies            map[string]bool
}

type NormalizetionRef struct {
	MaxTxTotal               int64
	MaxTxIn                  int64
	MaxTxOut                 int64
	MaxTxTotalVolumeUsdOnDay float64
	MaxTxInVolumeUsdOnDay    float64
	MaxTxOutVolumeUsdOnDay   float64
	MaxBalanceUsd            float64
}

type FetchingResult struct {
	Wallets map[string]Wallet
	Ref     NormalizetionRef
}

type NormalizedWallet struct {
	Address               string
	TxTotal               float64
	TxIn                  float64
	TxOut                 float64
	TxTotalVolumeUsdOnDay float64
	TxInVolumeUsdOnDay    float64
	TxOutVolumeUsdOnDay   float64
	BalanceUsd            float64
}

var ctx = context.Background()

func updateWalletBalance(wallet Wallet, client *eth.EthClient, redis *redis.RedisClient, blockTime uint64) Wallet {
	walletAddress := common.HexToAddress(wallet.Address)
	res := ethUsdBalance(walletAddress, client, redis, blockTime)

	for currency := range wallet.Currencies {
		if currency == "nil" {
			continue
		}
		if currency == eth.ETH {
			continue
		}
		//log.Debug().Msgf("currency: %s", currency)
		res += erc20Balance(walletAddress, client, redis, blockTime, currency)
	}

	//log.Debug().Msgf("addr: %s; bal: %f", walletAddress, res)
	wallet.BalanceUsd = res
	return wallet
}

func updateBalances(wallets map[string]Wallet, client *eth.EthClient, redis *redis.RedisClient, blockTime uint64) {
	for addr, wallet := range wallets {
		wallet = updateWalletBalance(wallet, client, redis, blockTime)
		wallets[addr] = wallet
	}
}

func checkTxStatus(client *ethclient.Client, txhash string, ctx context.Context) bool {
	reciept, err := client.TransactionReceipt(ctx, common.HexToHash(txhash))
	if err != nil {
		log.Fatal().Err(err).Msg("error getting reciept")
	}
	return reciept.Status == 1
}

func fetchWallets(
	redis *redis.RedisClient,
	client *eth.EthClient,
) *FetchingResult {
	wallets := make(map[string]Wallet)
	ref := NormalizetionRef{}

	block_number := START_BLOCK_NUMBER

	finishBlock, err := client.Client.BlockByNumber(ctx, big.NewInt(int64(FINISH_BLOCK_NUMBER)))
	if err != nil {
		log.Fatal().Err(err).Msg("error fetching finish block")
	}
	finishBlockTime := finishBlock.Time()
	//blockEthPriceUsd, err := eth.GetTokenPrice(finishBlockTime, redis, eth.ETH)
	if err != nil {
		log.Fatal().Err(err).Msg("error getting eth price for block")
	}

	for block_number <= FINISH_BLOCK_NUMBER {

		blockNumberStr := strconv.Itoa(block_number)
		log.Info().Msgf("block: %s", blockNumberStr)

		block, err := redis.GetBlock(&blockNumberStr)
		if err != nil {
			log.Fatal().Err(err).Msg("error getting block")
		}
		for _, tx := range strings.Split(*block, "\n") {
			if tx == "" {
				continue
			}
			vals := strings.Split(tx, ";")
			from := vals[0]
			//txhash := vals[1]
			to := vals[2]
			ethUsdOnDay := vals[4]
			ticker := vals[5]
			tickerUsdOnDay := vals[7]

			// check status is too expensive by time
			// let's cut 3 sigma deviations
			// isValid := checkTxStatus(client, txhash, ctx)
			// if !isValid {
			// 	continue
			//}
			if err != nil {
				log.Fatal().Err(err).Msg("error getting reciept")
			}
			totalTxUsd, err := strconv.ParseFloat(ethUsdOnDay, 64)
			if err != nil {
				log.Fatal().Err(err).Msg("error converting ethUsdOnDay to int")
			}

			if eth.IsCurrency(ticker) {
				erc20, err := decimal.NewFromString(tickerUsdOnDay)
				if err != nil {
					log.Fatal().Err(err).Msg("error converting ethUsdOnDay to int")
				}

				tokenPrice, err := eth.GetTokenPrice(finishBlockTime, redis, ticker)
				if err != nil {
					log.Error().Err(err).Msgf("failed to get token price %s", ticker)
					continue
				}

				erc20Usd := erc20.Mul(*tokenPrice)
				if erc20Usd.Cmp(MAX_USDT_LIMIT) >= 0 {
					continue
				}
				erc20UsdFloat, _ := erc20Usd.Float64()
				totalTxUsd += erc20UsdFloat
			}
			if totalTxUsd == 0 {
				// means we've not implemented this contract yet
				continue
			}
			if totalTxUsd < 0 || totalTxUsd > float64(math.MaxInt64) {
				continue
			}
			wFrom, fromExists := wallets[from]
			if !fromExists {
				wFrom = Wallet{
					Address:               from,
					TxTotal:               1,
					TxOut:                 1,
					TxTotalVolumeUsdOnDay: totalTxUsd,
					TxOutVolumeUsdOnDay:   totalTxUsd,
					Currencies:            map[string]bool{ticker: true},
				}
			} else {
				wFrom.TxTotal++
				wFrom.TxOut++
				wFrom.TxTotalVolumeUsdOnDay = totalTxUsd
				wFrom.TxOutVolumeUsdOnDay += totalTxUsd
				wFrom.Currencies[ticker] = true
				// wFrom2, _ := wallets[from]
				// if wFrom2.TxTotal < 2 {
				// 	log.Fatal().Msgf("wallet: %+v %+v", wFrom, wFrom2)
				// }
			}
			wallets[from] = wFrom
			ref = updateRef(ref, wFrom)

			wTo, toExists := wallets[to]
			if !toExists {
				wTo = Wallet{
					Address:               to,
					TxTotal:               1,
					TxIn:                  1,
					TxTotalVolumeUsdOnDay: totalTxUsd,
					TxInVolumeUsdOnDay:    totalTxUsd,
					Currencies:            map[string]bool{ticker: true},
				}
			} else {
				wTo.TxTotal++
				wTo.TxIn++
				wTo.TxTotalVolumeUsdOnDay += totalTxUsd
				wTo.TxInVolumeUsdOnDay += totalTxUsd
				wTo.Currencies[ticker] = true
			}
			wallets[to] = wTo
			ref = updateRef(ref, wTo)
		}
		block_number++
	}

	updateBalances(wallets, client, redis, finishBlockTime)

	return &FetchingResult{Wallets: wallets, Ref: ref}
}

func updateRef(ref NormalizetionRef, w Wallet) NormalizetionRef {
	if ref.MaxTxTotal < w.TxTotal {
		ref.MaxTxTotal = w.TxTotal
	}
	if ref.MaxTxIn < w.TxIn {
		ref.MaxTxIn = w.TxIn
	}
	if ref.MaxTxOut < w.TxOut {
		ref.MaxTxOut = w.TxOut
	}
	if ref.MaxTxTotalVolumeUsdOnDay < w.TxTotalVolumeUsdOnDay {
		ref.MaxTxTotalVolumeUsdOnDay = w.TxTotalVolumeUsdOnDay
	}
	if ref.MaxTxInVolumeUsdOnDay < w.TxInVolumeUsdOnDay {
		ref.MaxTxInVolumeUsdOnDay = w.TxInVolumeUsdOnDay
	}
	if ref.MaxTxOutVolumeUsdOnDay < w.TxOutVolumeUsdOnDay {
		ref.MaxTxOutVolumeUsdOnDay = w.TxOutVolumeUsdOnDay
	}
	if ref.MaxBalanceUsd < w.BalanceUsd {
		ref.MaxBalanceUsd = w.BalanceUsd
	}
	return ref
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	redis := redis.NewClient()
	client, err := eth.NewEthClient()
	if err != nil {
		log.Err(err).Msg("error connecting to Ethereum node")
		return
	}
	log.Info().Msg("Connected to Ethereum node")

	result := fetchWallets(redis, client)
	log.Info().Msgf("wallet: %+v", result.Ref)

	//nWallets := normalizeWalletsOld(*result)
	writeCsv(result.Wallets)
	//createClasters(nWallets)
}
