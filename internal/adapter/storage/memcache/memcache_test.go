package memcache

import (
	"context"
	"fmt"
	"gateway/internal/core/port/cache"
	"gateway/internal/core/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var (
	sleepTime = 2 * time.Second
	ctx       = context.Background()
)

func getConnection() cache.Memcache {
	newCache := NewMemcache()
	return newCache
}

func TestSetGetDelete(t *testing.T) {
	t.Parallel()

	engine := getConnection()
	time.Sleep(sleepTime)
	require.NotNil(t, engine)
	t.Run("set get delete", func(t *testing.T) {
		key := util.GenerateCacheKey("test", "container")
		value := []byte("test")
		err := engine.Set(ctx, key, value)
		require.NoError(t, err)
		value, err = engine.Get(ctx, key)
		require.NoError(t, err)
		require.Equal(t, "test", string(value))
		err = engine.Delete(ctx, key)
		require.NoError(t, err)
	})
	fmt.Println("Test set get delete done")
}
