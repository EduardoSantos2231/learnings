package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	value     any
	expiresAt time.Time
}

type Cache struct {
	data   map[string]cacheEntry
	mu     sync.RWMutex
	ticker *time.Ticker
	quit   chan struct{}
}

func NewCache(interval time.Duration) *Cache {
	result := &Cache{
		data:   make(map[string]cacheEntry),
		ticker: time.NewTicker(interval),
		quit:   make(chan struct{}),
	}
	go result.deleteExpiredKeys()
	return result
}

func (c *Cache) Stop() {
	c.ticker.Stop()
	close(c.quit)
}

func (c *Cache) deleteExpiredKeys() {
	for {
		select {
		case <-c.ticker.C:
			c.mu.Lock()
			for key, ce := range c.data {
				if time.Now().After(ce.expiresAt) {
					delete(c.data, key)
				}
			}
			c.mu.Unlock()
		case <-c.quit:
			return
		}

	}
}

func (c *Cache) Set(key string, value any, ttl time.Duration) {
	expireTime := time.Now().Add(ttl)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{
		value:     value,
		expiresAt: expireTime,
	}
}

func (c *Cache) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, exists := c.data[key]
	if !exists {
		return nil, false
	}
	if time.Now().After(entry.expiresAt) {
		return nil, false
	}
	return entry.value, true
}

func (c *Cache) Delete(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, exists := c.data[key]
	if !exists {
		return false
	}
	delete(c.data, key)
	return true
}
