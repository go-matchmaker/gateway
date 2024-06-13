//go:build wireinject
// +build wireinject

package app

import (
	"context"
	"gateway/internal/adapter/config"
	"gateway/internal/adapter/storage/memcache"
	adapter_http "gateway/internal/adapter/transport/http"
	"gateway/internal/core/port/cache"
	"gateway/internal/core/port/http"
	"github.com/google/wire"
	"go.uber.org/zap"
	"sync"
)

func InitApp(
	ctx context.Context,
	rw *sync.RWMutex,
	cfg *config.Container,
	gatewayLogger *zap.Logger,
) (*App, func(), error) {
	panic(wire.Build(
		New,
		httpServerFunc,
		memcache.MemCacheSet,
		memcache.MemCacheTTLSet,
	))
}

func httpServerFunc(
	ctx context.Context,
	Cfg *config.Container,
	gatewayLogger *zap.Logger,
	ttl cache.MemcacheTTL,
) (http.ServerMaker, func(), error) {
	httpServer := adapter_http.NewHTTPServer(ctx, Cfg, gatewayLogger, ttl)
	err := httpServer.Start(ctx)
	if err != nil {
		return nil, nil, err
	}
	return httpServer, func() { httpServer.Close(ctx) }, nil
}
