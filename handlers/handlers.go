package handlers

import (
	"fmt"
	"net/http"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {
	s := `{“action”: “add”, “x”: 2, “y”: 5, “answer”, 7, “cached”: false}`
	query := r.URL.Query()

	fmt.Printf("Query params: %v\n", query)
	fmt.Fprintf(w, s)
	fmt.Println("Endpoint Hit: /add")
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
