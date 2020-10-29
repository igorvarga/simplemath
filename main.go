package main

import (
	"github.com/igorvarga/teltechcodechallenge/cache"
	. "github.com/igorvarga/teltechcodechallenge/server"
	"log"
	"net/http"
	"time"
)

// TODO AddHandler REST API versioning
// TODO AddHandler godoc
// TODO AddHandler logging

// TODO Get parameters from env variables
var (
	addr     = ":80"
)

func main() {
	c := cache.NewCache(time.Minute, 5 * time.Minute)
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
