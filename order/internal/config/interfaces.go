package config

import (
	"github.com/IBM/sarama"
	"time"
)

type LoggerConfig interface {
	Level() string
	AsJson() bool
}

type OrderHTTPConfig interface {
	Address() string
	ReadTimeOut() time.Duration
}
type ExtDepConfig interface {
	InventoryAddress() string
	PaymentAddress() string
}

type PostgresConfig interface {
	URI() string
	DatabaseName() string
	MigrationsDir() string
}

type KafkaConfig interface {
	Brokers() []string
}

type OrderPaidProducerConfig interface {
	Topic() string
	Config() *sarama.Config
}

type OrderAssembledConsumerConfig interface {
	Topic() string
	GroupID() string
	Config() *sarama.Config
}
