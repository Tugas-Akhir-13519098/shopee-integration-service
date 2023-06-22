package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	KafkaHost                 string `envconfig:"KAFKA_HOST" default:"localhost"`
	KafkaPort                 string `envconfig:"KAFKA_PORT" default:"9092"`
	KafkaProductTopic         string `envconfig:"KAFKA_PRODUCT_TOPIC" default:"product"`
	KafkaProductConsumerGroup string `envconfig:"KAFKA_PRODUCT_CONSUMER_GROUP" default:"shopee-product-consumer-group"`

	ShopeeURL      string `envconfig:"SHOPEE_URL" default:"https://fb724b04-b7cc-47d0-bba0-21e93dda67b8.mock.pstmn.io/api/v2/product/"`
	OmnichannelURL string `envconfig:"OMNICHANNEL_URL" default:"http://localhost:8080/api/v1/product/marketplace/"`
	AdminToken     string `envconfig:"ADMIN_TOKEN" default:""`
}

// Get to get defined configuration
func Get() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)

	return cfg
}
