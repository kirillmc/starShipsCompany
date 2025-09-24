package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kirillmc/starShipsCompany/notification/internal/config/env"
)

type config struct {
	Logger                 LoggerConfig
	Kafka                  KafkaConfig
	OrderAssembledConsumer OrderAssembledConsumerConfig
	OrderPaidConsumer      OrderPaidConsumerConfig
	Telegram               TelegramConfig
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

	orderAssembledConsumer, err := env.NewOrderAssembledConsumerConfig()
	if err != nil {
		return err
	}

	orderPaidConsumer, err := env.NewOrderPaidConsumerConfig()
	if err != nil {
		return err
	}

	telegram, err := env.NewTelegramConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		Logger:                 loggerCfg,
		Kafka:                  kafka,
		OrderAssembledConsumer: orderAssembledConsumer,
		OrderPaidConsumer:      orderPaidConsumer,
		Telegram:               telegram,
	}

	return nil
}
