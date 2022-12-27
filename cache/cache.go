package cache

import (
	"sync"
	"time"
)

type CacheItem struct {
	value      interface{}
	expiration time.Time
}

type Cache struct {
	storage map[string]CacheItem
	sync    *sync.RWMutex
}

func New() Cache {
	cache := Cache{
		storage: make(map[string]CacheItem),
		sync:    new(sync.RWMutex),
	}
	go cacheCleaner(&cache)
	return cache
}

func (c Cache) Set(name string, value interface{}, ttl time.Duration) {
	c.sync.Lock()
	c.storage[name] = CacheItem{
		value:      value,
		expiration: time.Now().Add(ttl),
	}
	c.sync.Unlock()
}

func (c Cache) Get(name string) (interface{}, bool) {
	c.sync.RLock()
	defer c.sync.RUnlock()
	cacheItem, exists := c.storage[name]
	return cacheItem.value, exists
}

func (c Cache) Delete(name string) {
	c.sync.Lock()
	delete(c.storage, name)
	c.sync.Unlock()
}

func cacheCleaner(cache *Cache) {
	for {
		for name, value := range cache.storage {
			if isExpired(value.expiration) {
				cache.Delete(name)
			}
		}
	}
}

func isExpired(expiration time.Time) bool {
	if time.Now().After(expiration) {
		return true
	}
	return false
}
