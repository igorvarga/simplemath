package cache

import (
	"sync"
	"time"
)

type Cache interface {
	Load(key string) (it item, ok bool)
	Store(key string, it item)
}

type Item interface {
	Created() time.Time
}

type cache struct {
	sync.Mutex
	storage map[string]item
}

type item struct {
	payload *[]byte
	created time.Time
}

func (i item) Created() time.Time {
	panic("implement me")
}

func NewCache() Cache {
	return &cache{
		Mutex:   sync.Mutex{},
		storage: make(map[string]item),
	}
}

func NewItem(p *[]byte) Item {
	return &item{
		payload: p,
		created: time.Time{},
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