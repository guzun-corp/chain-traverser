package redis

import (
	"fmt"
	"math/big"

	"github.com/rs/zerolog/log"
)

func blocksByAddrKey(addr *string) string {
	return fmt.Sprintf("b%s:%s", DB_VERSION, *addr)
}

func (client RedisClient) GetAddressBlocks(addr *string) (*[]string, error) {
	key := blocksByAddrKey(addr)
	log.Debug().Msgf("getting blocks for %s", key)
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
	return fmt.Sprintf("tx%s:%s", DB_VERSION, *blockNumber)
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
	// log.Info().Msgf("add block key=%s", key)

	err := client.redis.Set(ctx, key, *blob, 0)
	if err.Err() != nil {
		log.Err(err.Err()).Msg("cant add block")
	}
}

func (client RedisClient) GetLastBlockNumber() (*int64, error) {
	key := fmt.Sprintf("meta:last_block%s", DB_VERSION)
	val, err := client.redis.Get(ctx, key).Int64()
	if err != nil {
		log.Err(err).Msg("Cant get last block number")
		return nil, err
	}
	return &val, nil
}

func (client RedisClient) UpdateLastBlockNumber(blockNumber *big.Int) error {
	key := fmt.Sprintf("meta:last_block%s", DB_VERSION)
	value := blockNumber.String()
	err := client.redis.Set(ctx, key, value, 0)
	if err.Err() != nil {
		log.Err(err.Err()).Msg("UpdateLastBlockNumber")
		return err.Err()
	}
	return nil
}
