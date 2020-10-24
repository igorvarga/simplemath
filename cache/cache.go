package cache

import (
	"sync"
	"time"
)

type Cache interface {
	Load(key string) (i Item, ok bool)
	Store(key string, value interface{}) (ok bool)
	ItemExpired(key string) (expired bool, ok bool)
}

type Item interface {
	Value() interface{}
	Expires() time.Time
}

type cache struct {
	sync.Mutex
	storage    map[string]Item
	expiration time.Duration
}

type item struct {
	value   interface{}
	expires time.Time
}

func (i *item) Value() interface{} {
	return i.value
}

func (i *item) Expires() time.Time {
	return i.expires
}

func (c *cache) ItemExpired(key string) (expired bool, ok bool) {
	c.Lock()

	defer c.Unlock()

	return c.itemExpired(key)
}

func (c *cache) itemExpired(key string) (expired bool, ok bool) {
	if i, ok := c.load(key); ok {
		return i.Expires().Before(time.Now()), true
	}

	return false, false
}

func NewCache(expiration time.Duration) Cache {
	return &cache{
		Mutex:      sync.Mutex{},
		storage:    make(map[string]Item),
		expiration: expiration,
	}
}

func (c *cache) Load(key string) (i Item, ok bool) {
	c.Lock()

	defer c.Unlock()

	return c.Load(key)
}

func (c *cache) load(key string) (i Item, ok bool) {
	i, ok = c.storage[key]

	return i, ok
}

func (c *cache) Store(key string, value interface{}) (ok bool) {
	c.Lock()

	defer c.Unlock()

	return c.store(key, value)
}

func (c *cache) store(key string, value interface{}) (ok bool) {
	expired, ok := c.itemExpired(key)

	if !expired && ok {
		return false
	}

	c.storage[key] = c.NewItem(value)

	return true
}

func (c *cache) NewItem(value interface{}) *item {
	return &item{
		value:   value,
		expires: time.Now().Add(c.expiration),
	}
}
