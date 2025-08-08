package pokecache

import (
	"sync"
	"time"
)



func NewCache(interval time.Duration) Cache {
	mux := &sync.RWMutex{}
	memory := map[string]cacheEntry{};
	newCache := Cache{
		memory: memory,
		mux: mux,
	}
	go newCache.readLoop(interval)
	return newCache
}


func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	c.memory[key] = cacheEntry{
		val: val,
		createdAt: time.Now(),
	}
	c.mux.Unlock()
}


func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	val, ok := c.memory[key]
	c.mux.RUnlock()
	return val.val, ok
}


func (c *Cache) readLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	
	for {
		<-ticker.C
		c.mux.Lock()
		for key, value := range c.memory {
			if time.Since(value.createdAt) >= interval {
				delete(c.memory, key)
			}
		}
		c.mux.Unlock()
	}
}
