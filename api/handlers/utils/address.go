package utils

import (
	"chain-traverser/api/handlers/schemas"
	"chain-traverser/internal/storage/redis"

	"github.com/rs/zerolog/log"
)

func AddressLabel(address string, redis *redis.RedisClient) string {
	// !TODO fetch labes from redis if they are exist
	return address[len(address)-8:]
}

func FetchAddress(address string, redis *redis.RedisClient) schemas.Node {
	cnt, err := redis.GetAddressTxNumber(&address)
	if err != nil {
		log.Err(err).Msgf("GetAddressTxNumber failed | address: %s", address)
	}
	labels, _ := redis.GetAddressLabels(&address)
	var primeLabel string
	var adType string
	if labels != nil {
		if labels.Prime != "" {
			primeLabel = labels.Prime
		} else {
			primeLabel = address[len(address)-8:]
		}
		adType = labels.Type
		if labels.Seconary != nil {
			for _, lbl := range *labels.Seconary {
				if lbl == "OFAC Sanctions Lists" {
					adType = "high_risk"
				}
			}
		}
	} else {
		primeLabel = address[len(address)-8:]
		adType = ""
	}

	return schemas.Node{Id: address, Label: primeLabel, Cnt: cnt, Picked: false, Type: adType}
}
