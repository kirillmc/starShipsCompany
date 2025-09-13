package env

import (
	"github.com/caarlos0/env/v11"
	"net"
)

type extDepEnvConfig struct {
	InventoryHost string `env:"INVENTORY_GRPC_HOST,required"`
	InventoryPort string `env:"INVENTORY_GRPC_PORT,required"`

	PaymentHost string `env:"PAYMENT_GRPC_HOST,required"`
	PaymentPort string `env:"PAYMENT_GRPC_PORT,required"`
}

type extDepConfig struct {
	raw extDepEnvConfig
}

func NewExtDepConfig() (*extDepConfig, error) {
	var raw extDepEnvConfig
	if err := env.Parse(&raw); err != nil {
		return nil, err
	}

	return &extDepConfig{raw: raw}, nil
}

func (cfg *extDepConfig) InventoryAddress() string {
	return net.JoinHostPort(cfg.raw.InventoryHost, cfg.raw.InventoryPort)
}

func (cfg *extDepConfig) PaymentAddress() string {
	return net.JoinHostPort(cfg.raw.PaymentHost, cfg.raw.PaymentPort)
}
