package main

import (
	"os"
	"time"

	"chain-traverser/internal/blockchain/eth"
	"chain-traverser/internal/config"
	"chain-traverser/internal/storage/redis"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const WORKER_TIMEOUT = 1 * time.Hour

func updatePriceData(redis *redis.RedisClient, priceClient *ccClient) error {
	start := time.Now()

	log.Info().Msg("Start pricing update...")

	for _, currency := range eth.CURRENCIES {

		priceData, err := priceClient.fetchPriceData(currency)
		if err != nil {
			return err
		}

		for _, data := range *priceData {
			err = redis.SetPriceData(data.Timestamp, data.PriceUSD, currency)
			if err != nil {
				return err
			}
			// log.Info().Msgf("price %s date %d", data.PriceUSD, data.Timestamp)
		}
		log.Info().Msgf("Pricing %s updated successful in %s", currency, time.Since(start))
	}
	return nil
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	cfg, cErr := config.NewConfig()
	if cErr != nil {
		log.Err(cErr).Msg("Error loading config")
		return
	}
	redis := redis.NewClient(&cfg.Redis)
	ccClient := NewCCClient(cfg.CryptoCompare.ApiKey)
	// stop := make(chan os.Signal, 1)
	// signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	for {
		err := updatePriceData(redis, ccClient)
		if err != nil {
			log.Err(err).Msg("Error updating pricing data")
		}
		log.Info().Msgf("Finished, waiting")

		time.Sleep(WORKER_TIMEOUT)

	}
}
