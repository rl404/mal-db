package cacher

import (
	"time"

	"github.com/rl404/mal-db/internal/config"
	"github.com/rl404/mal-db/internal/logger"
	"github.com/rl404/mal-plugin/cache/bigcache"
	"github.com/rl404/mal-plugin/cache/memcache"
	"github.com/rl404/mal-plugin/cache/nocache"
	"github.com/rl404/mal-plugin/cache/redis"
)

// Cacher is cache interface.
type Cacher interface {
	Get(key string, data interface{}) error
	Set(key string, data interface{}) error
	Delete(key string) error
	Close() error
}

// New to create new cacher. Support in-memory, redis, and memcache caching.
func New(l logger.Logger, cacheType string, address string, password string, expiredTime time.Duration) (Cacher, error) {
	if expiredTime <= 0 {
		expiredTime = 24 * time.Hour
	}

	switch cacheType {
	case config.InMemory:
		l.Debug("using in-memory cache: %v", expiredTime)
		return bigcache.New(expiredTime)
	case config.Redis:
		l.Debug("using redis cache: %v", expiredTime)
		return redis.New(address, password, expiredTime)
	case config.Memcache:
		l.Debug("using memcache cache: %v", expiredTime)
		return memcache.New(address, expiredTime)
	default:
		l.Warn("not using cache (recommended using in-memory cache)")
		return nocache.New()
	}
}
