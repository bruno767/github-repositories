package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCache_Put_Get(t *testing.T) {
	ttl, _ := time.ParseDuration("1s")
	cache := NewTTLCache(ttl, 1)

	assert.Nil(t, cache.Put("key", 12))
	assert.Nil(t, cache.Put("key", 14))

	cacheValue, getError := cache.Get("UnknownKey")
	assert.Nil(t, cacheValue)
	assert.Equal(t, ErrorNotFound, getError)

	cacheValue, getError = cache.Get("key")
	assert.Nil(t, getError)
	assert.Equal(t, 14, cacheValue)

	time.Sleep(ttl)
	cacheValue, getError = cache.Get("key")
	assert.Nil(t, cacheValue)
	assert.Equal(t, ErrorNotFound, getError)
}

