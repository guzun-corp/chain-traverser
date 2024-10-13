package redis

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

func pubAddrKey() string {
	return fmt.Sprintf("aq%s", DB_VERSION)
}

func (client RedisClient) SendAddress(addr string) {
	key := pubAddrKey()
	err := client.redisQueue.LPush(ctx, key, addr)
	if err.Err() != nil {
		log.Err(err.Err()).Msg("Cant publish address")
	}
	// log.Debug().Msgf("send addr %s", addr)

}
