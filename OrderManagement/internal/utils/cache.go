package utils

import (
	"time"

	cache "github.com/patrickmn/go-cache"
)

type DataCache struct {
	cache *cache.Cache
}

func NewDataCache() *DataCache {
	cache := cache.New(5*time.Minute, 10*time.Minute)
	return &DataCache{
		cache: cache,
	}
}

// set

func (d *DataCache) Set(key string, value interface{}, expiration time.Duration) {
	d.cache.Set(key, value, expiration)
}

func (d *DataCache) Get(key string) (interface{}, bool) {
	return d.cache.Get(key)
}

func (d *DataCache) Delete(key string) {
	d.cache.Delete(key)
}
