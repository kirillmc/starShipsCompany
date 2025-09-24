package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kirillmc/starShipsCompany/order/internal/app"
	"github.com/kirillmc/starShipsCompany/order/internal/config"
	"github.com/kirillmc/starShipsCompany/platform/pkg/closer"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	"go.uber.org/zap"
)

const (
	cfgPath        = "./deploy/compose/order/.env"
	gfShdwnTimeOut = 5 * time.Second
)

func main() {
	err := config.Load(cfgPath)
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	appCtx, appCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer appCancel()
	defer gracefulShutdown()

	closer.Configure(syscall.SIGINT, syscall.SIGTERM)

	a, err := app.New(appCtx)
	if err != nil {
		logger.Error(appCtx, "❌ Не удалось создать приложение", zap.Error(err))
		return
	}

	go func() {
		err = a.Run(appCtx)
		if err != nil {
			logger.Error(appCtx, "❌ Ошибка при работе приложения", zap.Error(err))
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func gracefulShutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), gfShdwnTimeOut)
	defer cancel()

	if err := closer.CloseAll(ctx); err != nil {
		logger.Error(ctx, "❌ Ошибка при завершении работы", zap.Error(err))
	}
}
