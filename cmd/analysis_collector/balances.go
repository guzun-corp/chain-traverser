package main

import (
	"chain-traverser/internal/blockchain/eth"
	"chain-traverser/internal/storage/redis"
	"context"
	"math"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

func erc20Balance(address common.Address, client *eth.EthClient, redis *redis.RedisClient, blockTime uint64, currency string) float64 {
	token := client.GetTokenByTicker(currency)
	if token == nil {
		log.Fatal().Msgf("error getting token %s", currency)
	}

	bal, err := token.Instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal().Msgf("failed to fetch balance address %s", address.String())
	}
	price, err := eth.GetTokenPrice(blockTime, redis, token.Ticker)
	if err != nil {
		log.Fatal().Msgf("error getting token price %s", token.Ticker)
	}
	decBal := decimal.NewFromBigInt(bal, 0).Div(decimal.NewFromInt(int64(math.Pow10(token.Denomination))))
	usd := decBal.Mul(*price)
	roundedUsd := usd.RoundBank(2)
	floatUsd, _ := roundedUsd.Float64()
	// log.Debug().Msgf("addr: %s; curr: %s, bal: %f", address.String(), token.Ticker, floatUsd)
	return floatUsd
}

func ethUsdBalance(address common.Address, client *eth.EthClient, redis *redis.RedisClient, blockTime uint64) float64 {
	// return usd balance
	ethUsdRate, err := eth.GetTokenPrice(blockTime, redis, eth.ETH)
	if err != nil {
		log.Fatal().Msg("error getting eth price")
	}
	ethBalance, err := client.Client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal().Err(err).Msg("error getting eth balance")
	}
	eth := decimal.NewFromBigInt(ethBalance, 0).Div(decimal.NewFromInt(WEI_IN_ETH))
	usd := eth.Mul(*ethUsdRate)
	roundedUsd := usd.RoundBank(2)
	floatUsd, _ := roundedUsd.Float64()
	return floatUsd
}
