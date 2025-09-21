package app

import (
	"context"
	"github.com/kirillmc/starShipsCompany/assembly/internal/config"
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
		if err := a.runConsumer(ctx); err != nil {
			errCh <- errors.Errorf("consumer crashed: %v", err)
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

func (a *App) runConsumer(ctx context.Context) error {
	logger.Info(ctx, "🚀 OrderPaid Kafka consumer running")

	err := a.diContainer.OrderPaidConsumerService().RunConsumer(ctx)
	if err != nil {
		return err
	}

	return nil
}
