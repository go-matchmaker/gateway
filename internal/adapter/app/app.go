package app

import (
	"context"
	"gateway/internal/adapter/config"
	adapter_http "gateway/internal/adapter/transport/http"
	"gateway/internal/core/port/cache"
	"gateway/internal/core/port/http"
	"sync"

	"go.uber.org/zap"
)

type App struct {
	rw               *sync.RWMutex
	Cfg              *config.Config
	GatewayLogger    *zap.Logger
	AuthLogger       *zap.Logger
	ManagementLogger *zap.Logger
	HTTP             http.ServerMaker
	MemCache         cache.Memcache
	MemCacheTTL      cache.MemcacheTTL
}

func New(
	rw *sync.RWMutex,
	cfg *config.Config,
	gatewayLogger *zap.Logger,
	authLogger *zap.Logger,
	managementLogger *zap.Logger,
	http http.ServerMaker,
	memCache cache.Memcache,
	memCacheTTL cache.MemcacheTTL,
) *App {
	return &App{
		rw:               rw,
		Cfg:              cfg,
		GatewayLogger:    gatewayLogger,
		AuthLogger:       authLogger,
		ManagementLogger: managementLogger,
		HTTP:             http,
		MemCache:         memCache,
		MemCacheTTL:      memCacheTTL,
	}
}

func httpServerFunc(
	ctx context.Context,
	Cfg *config.Config,
	gatewayLogger *zap.Logger,
	authLogger *zap.Logger,
	managementLogger *zap.Logger,
) (http.ServerMaker, func(), error) {
	httpServer := adapter_http.NewHTTPServer(ctx, Cfg, gatewayLogger, authLogger, managementLogger)
	err := httpServer.Start(ctx)
	if err != nil {
		return nil, nil, err
	}
	return httpServer, func() { httpServer.Close(ctx) }, nil
}

func (a *App) Run(ctx context.Context) {
	a.GatewayLogger.Info("RUNNER!")
}
