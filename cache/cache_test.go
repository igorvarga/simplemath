package cache

import (
	"testing"
	"time"
)

func TestCache_Expired(t *testing.T) {
	key := "key"

	c := NewCache(time.Duration(10))

	c.Store(key, nil)

	time.Sleep(time.Duration(20))

	expired, _ := c.ItemExpired(key)

	if !expired {
		t.Errorf("!ItemExpired() = %v, want true", expired)
	}
}
