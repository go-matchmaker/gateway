package main

import (
	"context"
	"gateway/internal/adapter/app"
	"gateway/internal/adapter/config"
	"gateway/internal/adapter/logger/auth_logger"
	"gateway/internal/adapter/logger/gateway_logger"
	"gateway/internal/adapter/logger/management_logger"
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

	gateLogger := gateway_logger.CreateLogger(cfg.Log.Level)
	defer gateLogger.Sync()
	authLogger := auth_logger.CreateLogger(cfg.Services.Auth.LogLevel)
	defer authLogger.Sync()
	managementLogger := management_logger.CreateLogger(cfg.Services.Management.LogLevel)
	defer managementLogger.Sync()

	cleanup := prepareApp(ctx, wg, rw, cfg, gateLogger, authLogger, managementLogger)
	zap.S().Info("⚡ Service name:", cfg.App.Name)
	<-ctx.Done()
	zap.S().Info("Context signal received, shutting down")
	wg.Wait()
	zap.S().Info("Waiting for all goroutines to finish")
	cleanup()
	zap.S().Info("Shutting down successfully")
}

func prepareApp(ctx context.Context, wg *sync.WaitGroup, rw *sync.RWMutex, cfg *config.Config, gatewayLogger, authLogger, managementLogger *zap.Logger) func() {
	var errMsg error
	a, cleanUp, errMsg := app.InitApp(ctx, wg, rw, cfg, gatewayLogger, authLogger, managementLogger)
	if errMsg != nil {
		zap.S().Error("failed init app", errMsg)
		<-ctx.Done()
	}
	a.Run(ctx)
	return cleanUp
}
