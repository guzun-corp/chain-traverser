package redis

import (
	"chain-traverser/internal/storage"
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
)

// number of transactions per address
func addrCntKey(addr *string) string {
	return fmt.Sprintf("c%s:%s", DB_VERSION, *addr)
}

func (client RedisClient) GetAddressTxNumber(addr *string) (int64, error) {
	key := addrCntKey(addr)
	val, err := client.redis.Get(ctx, key).Int64()
	if err != nil {
		// log.Err(err).Msg("Cant get counter")
	}
	// log.Debug().Msgf("got counter %d for %s", val, *addr)
	return val, nil

}

func (client RedisClient) UpdateCounters(transMap map[string]int64) {
	for address, count := range transMap {
		key := addrCntKey(&address)
		err := client.redis.IncrBy(ctx, key, count)
		if err.Err() != nil {
			log.Err(err.Err()).Msg("Cant increment counter")
		}
	}
}

// total amount of transactions per address
func addrTxAmountKey(addr *string) string {
	return fmt.Sprintf("a%s:%s", DB_VERSION, *addr)
}

func (client RedisClient) GetAddressTxAmount(addr *string) (int64, error) {
	key := addrTxAmountKey(addr)
	val, err := client.redisAnalytics.Get(ctx, key).Int64()
	if err != nil {
		log.Err(err).Msg("Cant get total tx amount")
	}
	log.Debug().Msgf("got counter %d for %s", val, *addr)
	return val, nil
}

func (client RedisClient) UpdateAddressTxAmount(transMap map[string]int64) {
	// TODO! rewrite to pipeline
	for address, count := range transMap {
		key := addrTxAmountKey(&address)
		err := client.redisAnalytics.IncrBy(ctx, key, count)
		if err.Err() != nil {
			log.Err(err.Err()).Msg("Cant increment total tx amount")
		}
	}
}

// addresses labels
func addrLabels(addr *string) string {
	return fmt.Sprintf("lbl%s:%s", DB_VERSION, *addr)
}

func (client RedisClient) GetAddressLabels(addr *string) (*storage.Labels, error) {
	key := addrLabels(addr)
	val, err := client.redisAnalytics.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var labels storage.Labels
	err = json.Unmarshal([]byte(val), &labels)
	if err != nil {
		log.Err(err).Msg("Failed to unmarshal JSON")
		return nil, err
	}

	return &labels, nil
}

// we update address's labels from python code
// func (client RedisClient) UpdateAddressLabels(transMap map[string]int64) {
// 	// TODO! rewrite to pipeline
// 	for address, count := range transMap {
// 		key := addrLabels(&address)
// 		err := client.redisAnalytics.Set(ctx, key, count)
// 		if err.Err() != nil {
// 			log.Err(err.Err()).Msg("Cant increment total tx amount")
// 		}
// 	}
// }
