/*
Copyright 2018-2026 Ruohang Feng <rh@vonng.com>
*/
package server

import (
	"sync"
	"time"
)

// ttlCache is a small concurrent response cache with lazy expiry.
// Keys embed the snapshot version, so a reload naturally invalidates entries;
// expired and overflow entries are dropped opportunistically.
type ttlCache struct {
	mu      sync.RWMutex
	ttl     time.Duration
	maxSize int
	items   map[string]cacheEntry
}

type cacheEntry struct {
	body []byte
	exp  time.Time
}

func newTTLCache(ttl time.Duration, maxSize int) *ttlCache {
	if ttl <= 0 {
		ttl = 5 * time.Minute
	}
	if maxSize <= 0 {
		maxSize = 1024
	}
	return &ttlCache{ttl: ttl, maxSize: maxSize, items: map[string]cacheEntry{}}
}

func (c *ttlCache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	e, ok := c.items[key]
	c.mu.RUnlock()
	if !ok {
		return nil, false
	}
	if time.Now().After(e.exp) {
		c.mu.Lock()
		delete(c.items, key)
		c.mu.Unlock()
		return nil, false
	}
	return e.body, true
}

func (c *ttlCache) Set(key string, body []byte) {
	now := time.Now()
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.items) >= c.maxSize {
		// drop expired first; if still full, drop arbitrary entries (map order)
		for k, e := range c.items {
			if now.After(e.exp) {
				delete(c.items, k)
			}
		}
		for k := range c.items {
			if len(c.items) < c.maxSize {
				break
			}
			delete(c.items, k)
		}
	}
	c.items[key] = cacheEntry{body: body, exp: now.Add(c.ttl)}
}
