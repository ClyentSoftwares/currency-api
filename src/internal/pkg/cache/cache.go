package cache

import (
	"sync"
	"time"
)

type Cache struct {
	mutex   sync.RWMutex
	data    map[string]interface{}
	expires map[string]time.Time
}

func NewCache() *Cache {
	return &Cache{
		data:    make(map[string]interface{}),
		expires: make(map[string]time.Time),
	}
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = value
	c.expires[key] = time.Now().Add(duration)
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	value, exists := c.data[key]
	if !exists || time.Now().After(c.expires[key]) {
		return nil, false
	}
	return value, true
}
