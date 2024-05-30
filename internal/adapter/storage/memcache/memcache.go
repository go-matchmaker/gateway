package memcache

import (
	"context"
	"errors"
	"gateway/internal/core/port/cache"
	"github.com/maypok86/otter"
)

var (
	ErrCacheSet                = errors.New("error while trying to set cache")
	ErrCacheGet                = errors.New("error while trying to get cache")
	_           cache.Memcache = (*memcache)(nil)
)

type memcache struct {
	cache *otter.Cache[string, []byte]
}

func NewMemcache() cache.Memcache {
	otterCache, err := otter.MustBuilder[string, []byte](10_000).
		CollectStats().
		Cost(func(key string, value []byte) uint32 {
			return 1
		}).
		CollectStats().
		Build()

	if err != nil {
		panic(err)
	}
	return &memcache{
		cache: &otterCache,
	}
}

func (d *memcache) Set(ctx context.Context, key string, value []byte) error {
	ok := d.cache.Set(key, value)
	if !ok {
		return ErrCacheSet

	}
	return nil
}

func (d *memcache) Get(ctx context.Context, key string) ([]byte, error) {
	value, ok := d.cache.Get(key)
	if !ok {
		return nil, ErrCacheGet
	}
	return value, nil
}

func (d *memcache) Delete(ctx context.Context, key string) error {
	d.cache.Delete(key)
	return nil
}
