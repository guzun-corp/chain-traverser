package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type CryptoCompareConfig struct {
	ApiKey string
}

type EthConfig struct {
	NodeUrl string `envconfig:"ETH_NODE_URL" default:"http://localhost:8545"`
}

// RedisConfig holds the configuration for the Redis client
type RedisConfig struct {
	Address  string        `envconfig:"REDIS_ADDRESS" default:"localhost:6379"`
	Password string        `envconfig:"REDIS_PASSWORD" default:""`
	PoolSize int           `envconfig:"REDIS_POOL_SIZE" default:"10"`
	Timeout  time.Duration `envconfig:"REDIS_TIMEOUT" default:"5s"`

	// app's variables:
	// for storing blocks
	MAIN_DB int `envconfig:"REDIS_MAIN_DB" default:"0"`
	// for storing statistics
	ANALYTICS_DB int `envconfig:"REDIS_ANALYTICS_DB" default:"1"`
	// for interprocess communication
	QUEUE_DB   int    `envconfig:"REDIS_QUEUE_DB" default:"2"`
	DB_VERSION string `envconfig:"REDIS_DB_VERSION" default:"1"`
}

type IndexerConfig struct {
	StartBlockNumber  int64 `envconfig:"START_BLOCK_NUMBER" default:"19050000"`
	FinishBlockNumber int64 `envconfig:"FINISH_BLOCK_NUMBER" default:"99999999"`
}

type ApiConfig struct {
	GraphSizeLimit int `envconfig:"API_GRAPH_SIZE_OUTPUT_LIMIT" default:"5000"`
}

type Config struct {
	Redis         RedisConfig
	Eth           EthConfig
	Indexer       IndexerConfig
	Api           ApiConfig
	CryptoCompare CryptoCompareConfig
}

// NewConfig creates and returns a new Config instance with values from environment variables
func NewConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
