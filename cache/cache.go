package cache

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Cache interface {
	Load(key string) (i Item, ok bool)
	Store(key string, value interface{})
	ItemExpired(key string) (expired bool, i Item)
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

func (c *cache) ItemExpired(key string) (expired bool, i Item) {
	c.Lock()

	defer c.Unlock()

	return c.itemExpired(key)
}

func (c *cache) itemExpired(key string) (expired bool, i Item) {
	if i, ok := c.loadNoSlide(key); ok {
		return i.Expires().Before(time.Now()), i
	}

	return false, nil
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

func NewCache(expiration time.Duration, sweepInterval time.Duration) Cache {
	return newCache(expiration, sweepInterval)
}

func newCache(expiration time.Duration, sweepInterval time.Duration) *cache {
	c := &cache{
		Mutex:      sync.Mutex{},
		storage:    make(map[string]Item),
		expiration: expiration,
		sweeper: sweeper{
			interval: sweepInterval,
			sweeping: false,
			started:  false,
			stop:     make(chan bool),
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

	if ok {
		// slide the expiry on access
		// TODO explore option of using pointers here
		i = c.newItem(i.Value())
		c.storage[key] = i
	}

	return i, ok
}

func (c *cache) loadNoSlide(key string) (i Item, ok bool) {
	i, ok = c.storage[key]

	return i, ok
}

func (c *cache) Store(key string, value interface{}) {
	c.Lock()

	c.store(key, value)

	c.Unlock()
}

func (c *cache) store(key string, value interface{}) {
	c.storage[key] = c.newItem(value)
}

func (c *cache) newItem(value interface{}) Item {
	return &item{
		value:   value,
		expires: time.Now().Add(c.expiration),
	}
}
