package cache

import (
	"sync"
	"time"
)

type Cache interface {
	Load(key string) (i Item, ok bool)
	Store(key string, value *[]byte)
	ItemExpired(key string) (expired bool, ok bool)
}

type Item interface {
	Value() *[]byte
	Expires() time.Time
}

type cache struct {
	sync.Mutex
	storage    map[string]Item
	expiration time.Duration
}

type item struct {
	value   *[]byte
	expires time.Time
}

func (i *item) Value() *[]byte {
	return i.value
}

func (i *item) Expires() time.Time {
	return i.expires
}

func (c *cache) ItemExpired(key string) (expired bool, ok bool) {
	if i, ok := c.Load(key); ok {
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

	i, ok = c.storage[key]

	return i, ok
}

func (c *cache) Store(key string, value *[]byte) {
	c.Lock()

	i := &item{
		value:   value,
		expires: time.Now().Add(c.expiration),
	}

	c.storage[key] = i

	c.Unlock()
}
