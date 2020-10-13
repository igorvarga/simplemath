package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)
// TODO Refactor JSON building
func AddHandler(w http.ResponseWriter, r *http.Request) {
	// TODO Extract x, y check to function
	x, err := strconv.ParseFloat(r.URL.Query().Get("x"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := fmt.Sprintf(`{"code": %v, "message": "%v"}`, 1, err.Error())
		fmt.Fprintf(w, response)
		return
	}

	y, err := strconv.ParseFloat(r.URL.Query().Get("y"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := fmt.Sprintf(`{"code": %v, "message": "%v"}`, 1, err.Error())
		fmt.Fprintf(w, response)
		return
	}

	answer := x + y

	response := fmt.Sprintf(`{"action": "add", "x": %v, "y": %v, "answer", %v, "cached": false}`, x, y, answer)

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, response)

}

func SubtractHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"result": true}`)
}

func DivideHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"result": true}`)
}

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"result": true}`)
}

/*
func extractParams(r *http.Request) (x float64, y float64, err error) {

}
 */