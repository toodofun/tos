package cacheutil

import (
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type CacheItem[T any] struct {
	Value      T
	Expiration *time.Timer
}

type CountdownCache[T any] struct {
	items map[string]*CacheItem[T]
	mutex sync.Mutex
	ttl   time.Duration
}

func NewCountdownCache[T any](ttl time.Duration) *CountdownCache[T] {
	return &CountdownCache[T]{
		items: make(map[string]*CacheItem[T]),
		ttl:   ttl,
	}
}

func (c *CountdownCache[T]) Set(key string, value T, callbacks ...func()) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if item, found := c.items[key]; found {
		item.Expiration.Stop()
	}

	expiration := time.AfterFunc(c.ttl, func() {
		c.mutex.Lock()
		defer c.mutex.Unlock()

		logrus.Debugf("cache expired: %s", key)

		for _, callback := range callbacks {
			callback()
		}

		delete(c.items, key)
	})

	c.items[key] = &CacheItem[T]{
		Value:      value,
		Expiration: expiration,
	}
}

func (c *CountdownCache[T]) Get(key string) (T, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if item, found := c.items[key]; found {
		item.Expiration.Stop()
		item.Expiration.Reset(c.ttl)
		logrus.Debugf("cache hit: %s", key)
		return item.Value, true
	} else {
		var res T
		return res, false
	}
}

func (c *CountdownCache[T]) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if item, found := c.items[key]; found {
		item.Expiration.Stop()
		logrus.Debugf("cache deleted: %s", key)
		delete(c.items, key)
	}
}
