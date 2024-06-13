package memcache

import (
	"context"
	"gateway/internal/core/port/cache"
	"github.com/google/wire"
	"github.com/maypok86/otter"
	"time"
)

var (
	_              cache.MemcacheTTL = (*memcacheTTL)(nil)
	MemCacheTTLSet                   = wire.NewSet(NewMemcacheTTL)
)

type memcacheTTL struct {
	cache *otter.CacheWithVariableTTL[string, []byte]
}

func NewMemcacheTTL() cache.MemcacheTTL {
	otterCache, err := otter.MustBuilder[string, []byte](10_000).
		CollectStats().
		Cost(func(key string, value []byte) uint32 {
			return 1
		}).
		WithVariableTTL().
		Build()
	if err != nil {
		panic(err)
	}

	return &memcacheTTL{
		cache: &otterCache,
	}
}

func (d *memcacheTTL) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	ok := d.cache.Set(key, value, ttl)
	if !ok {
		return ErrCacheSet

	}
	return nil
}

func (d *memcacheTTL) Get(ctx context.Context, key string) ([]byte, error) {
	value, ok := d.cache.Get(key)
	if !ok {
		return nil, ErrCacheGet
	}
	return value, nil
}

func (d *memcacheTTL) Delete(ctx context.Context, key string) error {
	d.cache.Delete(key)
	return nil
}
