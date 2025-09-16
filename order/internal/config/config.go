package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kirillmc/starShipsCompany/order/internal/config/env"
)

type config struct {
	Logger    LoggerConfig
	OrderHTTP OrderHTTPConfig
	Postgres  PostgresConfig
	ExtDep    ExtDepConfig
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

	orderHTTPCfg, err := env.NewOrderHTTPConfig()
	if err != nil {
		return err
	}

	extDepCfg, err := env.NewExtDepConfig()
	if err != nil {
		return err
	}

	postgresCfg, err := env.NewPostgresConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		Logger:    loggerCfg,
		OrderHTTP: orderHTTPCfg,
		Postgres:  postgresCfg,
		ExtDep:    extDepCfg,
	}

	return nil
}
