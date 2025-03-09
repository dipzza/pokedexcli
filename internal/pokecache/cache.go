package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu sync.Mutex
}

type cacheEntry struct {
	val   		[]byte
	createdAt time.Time
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		cache: map[string]cacheEntry{},
	}
	go cache.reapLoop(interval)

	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		val: val,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) (value []byte, found bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, found := c.cache[key]

	return entry.val, found
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for {
		<- ticker.C
		c.mu.Lock()

		for key, val := range c.cache {
			if time.Since(val.createdAt) > interval {
				delete(c.cache, key)
			}
		}

		c.mu.Unlock()
	}
}