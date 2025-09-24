package config

import "time"

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
