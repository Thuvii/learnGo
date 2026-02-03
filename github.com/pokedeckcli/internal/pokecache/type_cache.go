package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheMap map[string]cacheEntry
	cacheMu  *sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cacheMap: make(map[string]cacheEntry),
		cacheMu:  &sync.Mutex{},
		interval: interval,
	}
	go cache.reaploop()
	return cache

}

func (C *Cache) Add(key string, val []byte) {
	C.cacheMu.Lock()
	C.cacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	C.cacheMu.Unlock()
}

func (C *Cache) Get(key string) ([]byte, bool) {
	C.cacheMu.Lock()
	defer C.cacheMu.Unlock()
	value, exist := C.cacheMap[key]
	if exist {
		return value.val, exist
	}
	return nil, exist

}

func (C *Cache) reaploop() {
	ticker := time.NewTicker(C.interval)
	defer ticker.Stop()
	for range ticker.C {
		C.reap()
	}

}

func (C *Cache) reap() {
	C.cacheMu.Lock()
	for key, value := range C.cacheMap {
		if time.Since(value.createdAt) > C.interval {
			delete(C.cacheMap, key)
		}
	}
	C.cacheMu.Unlock()
}
