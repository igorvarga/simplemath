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

	log.Println("Starting server.")

	err := http.ListenAndServe(":80", nil)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Server started successfully.")
	}

}

func main() {
	handleRequests()
}
