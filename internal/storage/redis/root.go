package redis

import (
	"context"
	"fmt"
	"math/big"

	"github.com/rs/zerolog/log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisClient struct {
	redis *redis.Client
}

func NewClient() *RedisClient {
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6389",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	res := RedisClient{redis: redis}
	return &res
}

func addrCntKey(addr *string) string {
	return fmt.Sprintf("c:%s", *addr)
}

func (client RedisClient) GetAddressTxNumber(addr *string) (int64, error) {
	key := addrCntKey(addr)
	val, err := client.redis.Get(ctx, key).Int64()
	if err != nil {
		log.Err(err).Msg("Cant get counter")
	}
	log.Debug().Msgf("got counter %d for %s", val, *addr)
	return val, nil

}
func (client RedisClient) UpdateCounters(transMap map[string]int64) {
	// TODO! rewrite to pipeline
	for address, count := range transMap {
		key := addrCntKey(&address)
		err := client.redis.IncrBy(ctx, key, count)
		if err.Err() != nil {
			log.Err(err.Err()).Msg("Cant increment counter")
		}
	}
}

func blocksByAddrKey(addr *string) string {
	return fmt.Sprintf("b:%s", *addr)
}

func (client RedisClient) GetAddressBlocks(addr *string) (*[]string, error) {
	key := blocksByAddrKey(addr)
	log.Info().Msgf("getting blocks for %s", key)
	slice, err := client.redis.LRange(ctx, key, 0, 1000).Result()
	if err != nil {
		log.Err(err).Msg("Cant get block list")
		return nil, err
	}
	// log.Debug().Msgf("blocks %s", slice)
	log.Debug().Msgf("blocks %d", len(slice))
	return &slice, nil
}

func (client RedisClient) AppendBlockNumbers(transMap map[string]int64, blockNumberInt *big.Int) {
	// TODO! rewrite to pipeline
	blockNumber := blockNumberInt.String()
	for addr, _ := range transMap {
		key := blocksByAddrKey(&addr)
		err := client.redis.RPush(ctx, key, blockNumber)
		if err.Err() != nil {
			log.Err(err.Err()).Msg("Cant append block to wallet")
		}
	}
}

func trxByBlockKey(blockNumber *string) string {
	return fmt.Sprintf("tx:%s", *blockNumber)
}

func (client RedisClient) GetBlock(blockNumber *string) (*string, error) {
	key := trxByBlockKey(blockNumber)
	val, err := client.redis.Get(ctx, key).Result()
	if err != nil {
		log.Err(err).Msg("Cant get block")
		return nil, err
	}
	return &val, nil
}

func (client RedisClient) AddBlock(blockNumberInt *big.Int, blob *string) {
	blockNumber := blockNumberInt.String()
	key := trxByBlockKey(&blockNumber)
	err := client.redis.Set(ctx, key, *blob, 0)
	if err.Err() != nil {
		log.Err(err.Err()).Msg("cant add block")
	}
}
