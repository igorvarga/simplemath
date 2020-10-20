package main

import (
	"github.com/igorvarga/teltechcodechallenge/handler/std"
	"log"
	"net/http"
)

// TODO Add REST API versioning
// TODO Add godoc
// TODO Add logging

func handleRequests() {
	http.HandleFunc("/add", std.AddHandler)
	http.HandleFunc("/subtract", std.SubtractHandler)
	http.HandleFunc("/divide", std.DivideHandler)
	http.HandleFunc("/multiply", std.MultiplyHandler)

	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
