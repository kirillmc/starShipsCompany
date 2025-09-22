package app

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/kirillmc/starShipsCompany/notification/internal/config"
	"github.com/kirillmc/starShipsCompany/platform/pkg/closer"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type App struct {
	diContainer *diContainer
}

func New(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	errCh := make(chan error, 2)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		if err := a.runOrderPaidConsumer(ctx); err != nil {
			errCh <- errors.Errorf("OrderPaid consumer crashed: %v", err)
		}
	}()

	go func() {
		if err := a.runOrderAssembledConsumer(ctx); err != nil {
			errCh <- errors.Errorf("OrderAssembled consumer crashed: %v", err)
		}
	}()

	select {
	case <-ctx.Done():
		logger.Info(ctx, "Shutdown signal received")
	case err := <-errCh:
		logger.Error(ctx, "Component crashed, shutting down", zap.Error(err))
		cancel()
		<-ctx.Done()
		return err
	}

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initDI,
		a.initLogger,
		a.initCloser,
		a.initTelegramBot,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initDI(_ context.Context) error {
	a.diContainer = NewDIContainer()
	return nil
}

func (a *App) initLogger(_ context.Context) error {
	return logger.Init(config.AppConfig().Logger.Level(), config.AppConfig().Logger.AsJson())
}

func (a *App) initCloser(_ context.Context) error {
	closer.SetLogger(logger.Logger())
	return nil
}

func (a *App) initTelegramBot(ctx context.Context) error {
	telegramBot := a.diContainer.TelegramBot(ctx)

	telegramBot.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact,
		func(Ctx context.Context, b *bot.Bot, update *models.Update) {
			logger.Info(ctx, "chat id", zap.Int64("chat_id", update.Message.Chat.ID))

			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text: fmt.Sprintf(
					"🛸🛸🛸 %s, приветствуем в нашейм БЮРО ПО ПРОИЗВОДСТВУ КОСМИЧЕСКИХ КОРАБЛЕЙ!!!🛸🛸🛸",
					update.Message.From.FirstName),
			})
			if err != nil {
				logger.Error(ctx, "failed to send activation message", zap.Error(err))
			}
		})

	go func() {
		logger.Info(ctx, "🤖 Telegram bot started...")
		telegramBot.Start(ctx)
	}()

	return nil
}

func (a *App) runOrderPaidConsumer(ctx context.Context) error {
	logger.Info(ctx, "🚀 OrderPaid Kafka consumer running")

	err := a.diContainer.OrderPaidConsumerService(ctx).RunOrderPaidConsumer(ctx)
	if err != nil {
		return err
	}

	return nil
}
func (a *App) runOrderAssembledConsumer(ctx context.Context) error {
	logger.Info(ctx, "🚀 OrderAssembled Kafka consumer running")

	err := a.diContainer.OrderAssembledConsumerService(ctx).RunOrderAssembledConsumer(ctx)
	if err != nil {
		return err
	}

	return nil
}
