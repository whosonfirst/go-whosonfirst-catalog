package utils

import (
	gocache "github.com/patrickmn/go-cache"
	"time"
)

var cache *gocache.Cache

func init() {
	cache = gocache.New(30*time.Minute, 60*time.Minute)
}

func CacheGet(key string) (interface{}, bool) {
	return cache.Get(key)
}

func CacheSet(key string, value interface{}) {
	cache.Set(key, value, gocache.DefaultExpiration)
}
