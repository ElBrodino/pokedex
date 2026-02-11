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

func NewCache(interval time.Duration) Cache {
	c := Cache{
		//init map
		mux: &sync.Mutex{},
		//init mutex
		cacheMap: make(map[string]cacheEntry),
	}

	//start bg cleanup
	go c.reapLoop(interval)
	return c
}

func (c *Cache) reapLoop(interval time.Duration) {}
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
}

time.Now().Add(-interval)
