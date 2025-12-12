package weather

import (
	"sync"
	"time"
)

type Cache struct {
	mu        sync.RWMutex
	value     *Weather
	expiresAt time.Time
	ttl       time.Duration
}

func NewCache(ttl time.Duration) *Cache {
	return &Cache{
		ttl: ttl,
	}
}

func (c *Cache) Get() (*Weather, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.value == nil || time.Now().After(c.expiresAt) {
		return nil, false
	}

	return c.value, true
}

func (c *Cache) Set(w *Weather) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value = w
	c.expiresAt = time.Now().Add(c.ttl)
}
