package gin_cache_middle

import (
	"sync"
	"time"
)

type MockCache struct {
	cache map[string][]byte // lazy init
	sync.RWMutex
}

func (c *MockCache) Set(key string, value []byte, expiration time.Duration) error {
	c.Lock()
	defer c.Unlock()
	if c.cache == nil {
		c.cache = make(map[string][]byte)
	}
	c.cache[key] = value
	return nil
}

func (c *MockCache) Get(key string) ([]byte, bool) {
	c.RLock()
	defer c.RUnlock()
	if c.cache == nil {
		return []byte{}, false
	}
	value, ok := c.cache[key]
	return value, ok
}
