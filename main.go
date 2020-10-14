package main

import (
	h "github.com/igorvarga/teletchcodechallenge/handler"
	"log"
	"net/http"
)

// TODO Add REST API versioning
// TODO Add godoc

func handleRequests() {
	http.HandleFunc("/add", h.AddHandler)
	http.HandleFunc("/subtract", h.SubtractHandler)
	http.HandleFunc("/divide", h.DivideHandler)
	http.HandleFunc("/multiply", h.MultiplyHandler)

	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
