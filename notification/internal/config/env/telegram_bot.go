package env

import "github.com/caarlos0/env/v11"

type telegramEnvConfig struct {
	TelegramBotToken string `env:"TELEGRAM_BOT_TOKEN,required"`
}

type telegramConfig struct {
	raw telegramEnvConfig
}

func NewTelegramConfig() (*telegramConfig, error) {
	var raw telegramEnvConfig
	if err := env.Parse(&raw); err != nil {
		return nil, err
	}

	return &telegramConfig{raw: raw}, nil
}

func (cfg *telegramConfig) Token() string {
	return cfg.raw.TelegramBotToken
}
