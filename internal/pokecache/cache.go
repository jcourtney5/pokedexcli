package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries  map[string]CacheEntry
	lock     sync.RWMutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries:  make(map[string]CacheEntry),
		lock:     sync.RWMutex{},
		interval: interval,
	}

	go cache.reapLoop()

	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.lock.Lock()
	defer cache.lock.Unlock()
	cache.entries[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.lock.RLock()
	defer cache.lock.RUnlock()
	entry, ok := cache.entries[key]
	if ok {
		return entry.val, ok
	} else {
		return nil, false
	}
}

func (cache *Cache) reapLoop() {
	ticker := time.NewTicker(cache.interval)

	defer ticker.Stop()

	for t := range ticker.C {
		//fmt.Println("Running reapLoop() at:", t.Format(time.RFC3339))

		cache.lock.Lock()

		for key, cacheEntry := range cache.entries {
			if t.Sub(cacheEntry.createdAt) > cache.interval {
				fmt.Printf("Deleting cache entry %s\n", key)
				delete(cache.entries, key)
			}
		}

		cache.lock.Unlock()
	}
}
