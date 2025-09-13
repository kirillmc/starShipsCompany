package config

import (
	"github.com/kirillmc/starShipsCompany/inventory/internal/config/env"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	Logger        LoggerConfig
	InventoryGRPC InventoryGRPCConfig
	Mongo         MongoConfig
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

	inventoryGRPCCfg, err := env.NewInventoryGRPCConfig()
	if err != nil {
		return err
	}

	mongoCfg, err := env.NewMongoConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		Logger:        loggerCfg,
		InventoryGRPC: inventoryGRPCCfg,
		Mongo:         mongoCfg,
	}

	return nil
}
