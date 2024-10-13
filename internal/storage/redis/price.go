package redis

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

func priceDataKey(timestamp int64, currency string) string {
	currency = strings.ToLower(currency)

	return fmt.Sprintf("%s:price%s:%d", currency, DB_VERSION, timestamp)
}

func (client *RedisClient) SetPriceData(timestamp int64, priceUSD, currency string) error {
	key := priceDataKey(timestamp, currency)
	err := client.redis.Set(ctx, key, priceUSD, 0)
	// log.Info().Msgf("key %s", key)
	if err.Err() != nil {
		log.Err(err.Err()).Msg("SetPriceData")
		return err.Err()
	}

	return nil
}

func (client *RedisClient) GetPrice(timestamp int64, currency string) (*decimal.Decimal, error) {
	key := priceDataKey(timestamp, currency)

	val, err := client.redis.Get(ctx, key).Result()
	if err != nil {
		log.Err(err).Msgf("GetPrice fetching %d %s", timestamp, currency)
		return nil, err
	}

	// log.Info().Msgf("getch key %s |%s|", key, val)

	price, err := decimal.NewFromString(val)

	if err != nil {
		log.Err(err).Msg("GetPrice type casting")
		return nil, err
	}

	return &price, nil
}
