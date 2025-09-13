package main

import (
	"context"
	"fmt"
	"github.com/kirillmc/starShipsCompany/payment/internal/app"
	"github.com/kirillmc/starShipsCompany/payment/internal/config"
	"github.com/kirillmc/starShipsCompany/platform/pkg/closer"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	"go.uber.org/zap"
	"os/signal"
	"syscall"
	"time"
)

const (
	cfgPath        = "./deploy/compose/payment/.env"
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

	err = a.Run(appCtx)
	if err != nil {
		logger.Error(appCtx, "❌ Ошибка при работе приложения", zap.Error(err))
		return
	}
}

func gracefulShutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), gfShdwnTimeOut)
	defer cancel()

	if err := closer.CloseAll(ctx); err != nil {
		logger.Error(ctx, "❌ Ошибка при завершении работы", zap.Error(err))
	}
}
