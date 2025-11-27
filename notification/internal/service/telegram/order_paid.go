package telegram

import (
	"bytes"
	"context"
	"embed"
	"html/template"

	"github.com/kirillmc/starShipsCompany/notification/internal/model"
)

//go:embed templates/paid_notification.tmpl
var orderPaidFS embed.FS

type orderPaidData struct {
	EventUUID       string
	OrderUUID       string
	UserUUID        string
	PaymentMethod   string
	TransactionUUID string
}

var orderPaidTemplate = template.Must(template.ParseFS(orderPaidFS, "templates/paid_notification.tmpl"))

func (s *service) SendOrderPaidNotification(ctx context.Context, orderPaidInfo model.OrderPaidEvent) error {
	message, err := s.buildOrderPaidMessage(orderPaidInfo)
	if err != nil {
		return err
	}

	err = s.telegramClient.SendMessage(ctx, chatID, message)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) buildOrderPaidMessage(orderPaidInfo model.OrderPaidEvent) (string, error) {
	data := orderPaidData{
		EventUUID:       orderPaidInfo.EventUUID,
		OrderUUID:       orderPaidInfo.OrderUUID,
		UserUUID:        orderPaidInfo.UserUUID,
		PaymentMethod:   orderPaidInfo.PaymentMethod,
		TransactionUUID: orderPaidInfo.TransactionUUID,
	}

	var buf bytes.Buffer
	err := orderPaidTemplate.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
