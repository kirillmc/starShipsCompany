package config

import (
	"github.com/kirillmc/starShipsCompany/payment/internal/config/env"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	Logger      LoggerConfig
	PaymentGRPC PaymentGRPCConfig
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

	paymentGRPCCfg, err := env.NewPaymentGRPCConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		Logger:      loggerCfg,
		PaymentGRPC: paymentGRPCCfg,
	}

	return nil
}
