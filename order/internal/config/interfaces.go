package config

type LoggerConfig interface {
	Level() string
	AsJson() bool
}

type OrderHTTPConfig interface {
	Address() string
}
type ExtDepConfig interface {
	InventoryAddress() string
	PaymentAddress() string
}

type PostgresConfig interface {
	URI() string
	DatabaseName() string
}
