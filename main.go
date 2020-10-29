package main

import (
	"github.com/igorvarga/teltechcodechallenge/cache"
	. "github.com/igorvarga/teltechcodechallenge/server"
	. "github.com/igorvarga/teltechcodechallenge/util"
	"log"
	"net/http"
	"time"
)

// TODO AddHandler REST API versioning
// TODO AddHandler godoc
// TODO AddHandler logging

var (
	addr          = GetEnv("SM_ADDR", ":8080")
	expiration    = GetEnvInt64("SM_CACHE_EXPIRATION", 60)
	sweepinterval = GetEnvInt64("SM_CACHE_SWEEPINTERVAL", 5)
)

func main() {

	c := cache.NewCache(time.Duration(expiration)*time.Second, time.Duration(sweepinterval)*time.Second)
	s := NewSimpleMathServer(c)

	http.HandleFunc("/add", s.CacheMiddleware(s.AddHandler))
	http.HandleFunc("/subtract", s.CacheMiddleware(s.SubtractHandler))
	http.HandleFunc("/divide", s.CacheMiddleware(s.DivideHandler))
	http.HandleFunc("/multiply", s.CacheMiddleware(s.MultiplyHandler))

	log.Printf("Server started on %v", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

