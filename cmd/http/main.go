package main

import (
	"context"
	"gateway/internal/adapter/app"
	"gateway/internal/adapter/config"
	"gateway/internal/adapter/logger"
	"gateway/internal/core/util"
	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
	"os/signal"
	"sync"
)

func main() {
	_, err := maxprocs.Set()
	if err != nil {
		panic("failed set max procs")
	}
	ctx, cancel := signal.NotifyContext(context.Background(), util.InterruptSignals...)
	defer cancel()
	wg := new(sync.WaitGroup)
	rw := new(sync.RWMutex)
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("failed get config: " + err.Error())
	}

	gateLogger := logger.CreateLogger(cfg.Log.Level)
	defer gateLogger.Sync()

	cleanup := prepareApp(ctx, wg, rw, cfg, gateLogger)
	zap.S().Info("⚡ Service name:", cfg.App.Name)
	<-ctx.Done()
	zap.S().Info("Context signal received, shutting down")
	wg.Wait()
	zap.S().Info("Waiting for all goroutines to finish")
	cleanup()
	zap.S().Info("Shutting down successfully")
}

func prepareApp(ctx context.Context, wg *sync.WaitGroup, rw *sync.RWMutex, cfg *config.Container, gatewayLogger *zap.Logger) func() {
	var errMsg error
	a, cleanUp, errMsg := app.InitApp(ctx, rw, cfg, gatewayLogger)
	if errMsg != nil {
		zap.S().Error("failed init app", errMsg)
		<-ctx.Done()
	}
	a.Run(ctx)
	return cleanUp
}
