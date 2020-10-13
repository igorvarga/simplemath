package handlers

import (
	"fmt"
	"net/http"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"result": true}`)
	fmt.Println("Endpoint Hit: /addHandler")
}

func SubtractHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"result": true}`)
	fmt.Println("Endpoint Hit: /subtract")
}

func DivideHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"result": true}`)
	fmt.Println("Endpoint Hit: /divide")
}

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"result": true}`)
	fmt.Println("Endpoint Hit: /multiply")
}
