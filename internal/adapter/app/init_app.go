package app

import (
	"context"
	"gateway/internal/adapter/config"
	"gateway/internal/adapter/storage/memcache"
	"go.uber.org/zap"
	"sync"
)

func InitApp(ctx context.Context, wg *sync.WaitGroup, rw *sync.RWMutex, Cfg *config.Config,
	gatewayLogger *zap.Logger) (*App, func(), error) {
	cacheMemcache := memcache.NewMemcache()
	memcacheTTL := memcache.NewMemcacheTTL()
	serverMaker, cleanup, err := httpServerFunc(ctx, Cfg, gatewayLogger, memcacheTTL)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	app := New(rw, Cfg, gatewayLogger, serverMaker, cacheMemcache, memcacheTTL)
	return app, func() {
		cleanup()
	}, nil
}
