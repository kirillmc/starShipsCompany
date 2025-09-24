package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kirillmc/starShipsCompany/assembly/internal/config/env"
)

type config struct {
	Logger                 LoggerConfig
	Kafka                  KafkaConfig
	OrderAssembledProducer OrderAssembledProducerConfig
	OrderPaidConsumer      OrderPaidConsumerConfig
}

var appConfig *config

func AppConfig() *config {
	return appConfig
}

func Load(path ...string) error {
	err := godotenv.Load(path...)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	loggerCfg, err := env.NewLoggerConfig()
	if err != nil {
		return err
	}

	kafka, err := env.NewKafkaConfig()
	if err != nil {
		return err
	}

	orderAssembledProducer, err := env.NewOrderAssembledProducerConfig()
	if err != nil {
		return err
	}

	orderPaidConsumer, err := env.NewOrderPaidConsumerConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		Logger:                 loggerCfg,
		Kafka:                  kafka,
		OrderAssembledProducer: orderAssembledProducer,
		OrderPaidConsumer:      orderPaidConsumer,
	}

	return nil
}
