package pokecache

import (
	"sync"
	"time"
)

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
			c.Mu.Lock()

			for k, v := range c.Dict {
				if time.Since(v.CreatedAt) > interval {
					delete(c.Dict, k)
			}	
		}
		c.Mu.Unlock()
	}
}


func (c *Cache) Get (key string) ([]byte, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	if data, ok := c.Dict[key]; ok {
		return data.Val, true	
	}
	return nil, false
}


func (c *Cache) Add (key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	if _, ok := c.Dict[key]; !ok {
		c.Dict[key] = CacheEntry{
			CreatedAt: time.Now(),
			Val: val,
		}
	}
	// else cache is already filled?
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache {
		Dict: make(map[string]CacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

type Cache struct {
	Mu sync.Mutex
	Dict map[string]CacheEntry
}

type CacheEntry struct {
	CreatedAt time.Time
	Val []byte
}