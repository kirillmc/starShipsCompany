package telegram

import "github.com/kirillmc/starShipsCompany/notification/internal/client/http"

const chatID = 320790952

type service struct {
	telegramClient http.TelegramClient
}

func NewService(telegramClient http.TelegramClient) *service {
	return &service{
		telegramClient: telegramClient,
	}
}
