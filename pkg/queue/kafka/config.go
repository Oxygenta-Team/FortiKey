package kafka

import "time"

type producer struct {
	ProducersAmount int    `yaml:"producersAmount"`
	Topic           string `yaml:"topic"`
}

type consumer struct {
	ConsumersAmount   int           `yaml:"consumersAmount"`
	Topics            []string      `yaml:"topics"`
	Group             string        `yaml:"group"`
	HeartbeatInterval time.Duration `yaml:"heartbeatInterval"`
}

type Config struct {
	Brokers  []string `yaml:"brokers"`
	Producer producer `yaml:"producer"`
	Consumer consumer `yaml:"consumer"`
}
