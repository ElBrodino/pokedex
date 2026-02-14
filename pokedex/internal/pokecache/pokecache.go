package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mux      *sync.Mutex
	cacheMap map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cacheMap[key] = cacheEntry{createdAt: time.Now().UTC(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, ok := c.cacheMap[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		//init map
		mux: &sync.Mutex{},
		//init mutex
		cacheMap: make(map[string]cacheEntry),
	}

	//start bg cleanup
	go c.reapLoop(interval)
	return c
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	cutoff := time.Now().Add(-interval)

	c.mux.Lock()
	defer c.mux.Unlock()

	for key, entry := range c.cacheMap {
		if entry.createdAt.Before(cutoff) {
			delete(c.cacheMap, key)
		}
	}
}
