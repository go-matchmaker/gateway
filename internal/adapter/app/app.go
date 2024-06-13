package app

import (
	"context"
	"gateway/internal/adapter/config"
	"gateway/internal/core/port/cache"
	"gateway/internal/core/port/http"
	"sync"

	"go.uber.org/zap"
)

type App struct {
	rw            *sync.RWMutex
	Cfg           *config.Container
	GatewayLogger *zap.Logger
	HTTP          http.ServerMaker
	MemCache      cache.Memcache
	MemCacheTTL   cache.MemcacheTTL
}

func New(
	rw *sync.RWMutex,
	cfg *config.Container,
	gatewayLogger *zap.Logger,
	http http.ServerMaker,
	memCache cache.Memcache,
	memCacheTTL cache.MemcacheTTL,
) *App {
	return &App{
		rw:            rw,
		Cfg:           cfg,
		GatewayLogger: gatewayLogger,
		HTTP:          http,
		MemCache:      memCache,
		MemCacheTTL:   memCacheTTL,
	}
}

func (a *App) Run(ctx context.Context) {
	a.GatewayLogger.Info("RUNNER!")
}
