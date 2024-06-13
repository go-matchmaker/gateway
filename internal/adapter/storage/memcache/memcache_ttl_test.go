package memcache

import (
	"fmt"
	"gateway/internal/core/port/cache"
	"gateway/internal/core/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func getConnectionTTL() cache.MemcacheTTL {
	newCache := NewMemcacheTTL()
	return newCache
}

func TestSetGetDeleteWithTTL(t *testing.T) {
	t.Parallel()

	engine := getConnectionTTL()
	time.Sleep(sleepTime)
	require.NotNil(t, engine)
	t.Run("set get delete", func(t *testing.T) {
		key := util.GenerateCacheKey("test", "container")
		value := []byte("test")
		err := engine.Set(ctx, key, value, 10*time.Second)
		require.NoError(t, err)
		value, err = engine.Get(ctx, key)
		require.NoError(t, err)
		require.Equal(t, "test", string(value))
		err = engine.Delete(ctx, key)
		require.NoError(t, err)
		err = engine.Set(ctx, key, value, 1*time.Second)
		require.NoError(t, err)
		time.Sleep(2 * time.Second)
		value, err = engine.Get(ctx, key)
		require.Error(t, err)
		require.Nil(t, value)
	})
	fmt.Println("Test set get delete done")
}
