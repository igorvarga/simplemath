package cache

import (
	"fmt"
	"log"
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

type sweeper struct {
	interval time.Duration
	ticker   *time.Ticker
	sweeping bool
	started  bool
	stop     chan bool
}

type cache struct {
	sync.Mutex
	storage    map[string]Item
	expiration time.Duration
	sweeper    sweeper
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

func (s *sweeper) start(c *cache) {
	if s.started {
		log.Print("Sweeper already started, returning.")
		return
	}

	s.started = true

	s.ticker = time.NewTicker(s.interval)

	go func() {
		for {
			select {
			case t := <-s.ticker.C:
				{
					fmt.Println("Running sweep at", t)
					c.sweep()
				}
			case <-s.stop:
				{
					s.ticker.Stop()
					return
				}
			}
		}
	}()
}

func (c *cache) sweep() {
	if c.sweeper.sweeping {
		log.Print("Sweeping in progress, returning.")
		return
	}

	c.Lock()

	c.sweeper.sweeping = true

	for key, item := range c.storage {
		if item.Expires().Before(time.Now()) {
			log.Printf("Evicting expired key %v from the cache.\n", key)
			delete(c.storage, key)
		}
	}

	c.sweeper.sweeping = false

	c.Unlock()

}

func NewCache(expiration time.Duration, interval time.Duration) Cache {
	return newCache(expiration, interval)
}

func newCache(expiration time.Duration, interval time.Duration) *cache {
	c := &cache{
		Mutex:      sync.Mutex{},
		storage:    make(map[string]Item),
		expiration: expiration,
		sweeper: sweeper{
			interval: interval,
			sweeping: false,
			started:  false,
			stop: make(chan bool),
		},
	}

	c.sweeper.start(c)

	return c
}

func (c *cache) Load(key string) (i Item, ok bool) {
	c.Lock()

	defer c.Unlock()

	return c.load(key)
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

	c.storage[key] = c.newItem(value)

	return true
}

func (c *cache) newItem(value interface{}) *item {
	return &item{
		value:   value,
		expires: time.Now().Add(c.expiration),
	}
}
