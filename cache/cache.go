package cache

import (
	"sync"
	"time"
)

type cache struct {
	sync.Mutex
	storage map[string]item
}

type item struct {
	payload *[]byte
	created time.Time
}

func NewCache() cache {
	return cache{
		Mutex:   sync.Mutex{},
		storage: make(map[string]item),
	}
}

func (c *cache) Load(key string) (it item, ok bool) {
	c.Lock()

	defer c.Unlock()

	it, ok = c.storage[key]

	return it, ok
}

func (c *cache) Store(key string, it item) {
	c.Lock()

	c.storage[key] = it

	c.Unlock()
}
