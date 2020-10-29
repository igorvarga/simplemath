package cache

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type TCache struct {
	*cache
}

func (tc *TCache) StopSweeper() {
	tc.sweeper.stop <- true
}

func NewTCache(expiration time.Duration, interval time.Duration) TCache {
	c := TCache{
		cache: newCache(expiration, interval),
	}

	return c
}

func TestCache_Expired(t *testing.T) {
	key := fmt.Sprint(rand.Int())

	c := NewTCache(200 * time.Millisecond, time.Second)

	c.Store(key, nil)

	time.Sleep(300 * time.Millisecond)

	expired, _ := c.ItemExpired(key)

	if !expired {
		t.Errorf("!ItemExpired() = %v, want true", expired)
	}

	c.StopSweeper()
}

func TestCache_Slide(t *testing.T) {
	key := fmt.Sprint(rand.Int())

	c := NewTCache(500 * time.Millisecond, time.Second)

	c.Store(key, nil)

	expired, _ := c.ItemExpired(key)

	fmt.Printf("Expired %v \n", expired)

	time.Sleep(250 * time.Millisecond)

	_, _ = c.Load(key)

	time.Sleep(400 * time.Millisecond)

	expired, _ = c.ItemExpired(key)

	if expired {
		t.Errorf("ItemExpired() = %v, want false", expired)
	}

	c.StopSweeper()
}

func TestCache_Eviction(t *testing.T) {
	c := NewTCache(time.Duration(10), 200 * time.Millisecond)

	key := fmt.Sprint(rand.Int())

	c.Store(key, nil)

	key1 := fmt.Sprint(rand.Int())

	c.Store(key1, nil)

	time.Sleep(300 * time.Millisecond)

	size := len(c.storage)

	if size != 0 {
		t.Errorf("Map size = %v, want 0", size)
	}

	c.StopSweeper()
}

func TestCache_SweepRunning(t *testing.T) {
	c := NewTCache(time.Duration(10), time.Second)

	c.cache.sweeper.sweeping = true

	c.sweeper.start(c.cache)

	key := fmt.Sprint(rand.Int())

	c.Store(key, nil)

	key1 := fmt.Sprint(rand.Int())

	c.Store(key1, nil)

	time.Sleep(2 * time.Second)

	size := len(c.storage)

	if size != 2 {
		t.Errorf("Map size = %v, want 2", size)
	}

	c.StopSweeper()
}