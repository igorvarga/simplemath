package cache

import "time"

var Cache = make(map[string]Item)

type Item struct {
	Payload *[]byte
	Created time.Time
}