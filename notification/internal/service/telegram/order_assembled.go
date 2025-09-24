package telegram

import (
	"bytes"
	"context"
	"embed"
	"html/template"

	"github.com/kirillmc/starShipsCompany/notification/internal/model"
)

//go:embed templates/assembled_notification.tmpl
var orderAssembledFS embed.FS

type orderAssembledData struct {
	EventUUID    string
	OrderUUID    string
	UserUUID     string
	BuildTimeSec int64
}

var orderAssembledTemplate = template.Must(template.ParseFS(orderAssembledFS, "templates/assembled_notification.tmpl"))

func (s *service) SendOrderAssembledNotification(ctx context.Context, orderAssembledInfo model.OrderAssembledEvent) error {
	message, err := s.buildOrderAssembledMessage(orderAssembledInfo)
	if err != nil {
		return err
	}

	err = s.telegramClient.SendMessage(ctx, chatID, message)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) buildOrderAssembledMessage(orderAssembledInfo model.OrderAssembledEvent) (string, error) {
	data := orderAssembledData{
		EventUUID:    orderAssembledInfo.EventUUID,
		OrderUUID:    orderAssembledInfo.OrderUUID,
		UserUUID:     orderAssembledInfo.UserUUID,
		BuildTimeSec: orderAssembledInfo.BuildTimeSec,
	}

	var buf bytes.Buffer
	err := orderAssembledTemplate.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
