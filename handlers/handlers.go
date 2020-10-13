package handlers

import (
	"fmt"
	sm "github.com/igorvarga/teletchcodechallenge/simplemath"
	"net/http"
	"strconv"
)

// TODO Refactor JSON building
func AddHandler(w http.ResponseWriter, r *http.Request) {
	x, y, err := extractParams(w, r)
	if err != nil {
		return
	}

	answer := sm.Add(x, y)

	response := fmt.Sprintf(`{"action": "add", "x": %v, "y": %v, "answer", %v, "cached": false}`, x, y, answer)

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, response)
}

func SubtractHandler(w http.ResponseWriter, r *http.Request) {
	x, y, err := extractParams(w, r)
	if err != nil {
		return
	}

	answer := sm.Subtract(x, y)

	response := fmt.Sprintf(`{"action": "subtract", "x": %v, "y": %v, "answer", %v, "cached": false}`, x, y, answer)

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, response)
}

func DivideHandler(w http.ResponseWriter, r *http.Request) {
	x, y, err := extractParams(w, r)
	if err != nil {
		return
	}

	answer := sm.Divide(x, y)

	response := fmt.Sprintf(`{"action": "divide", "x": %v, "y": %v, "answer", %v, "cached": false}`, x, y, answer)

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, response)
}

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	x, y, err := extractParams(w, r)
	if err != nil {
		return
	}

	answer := sm.Multiply(x, y)

	response := fmt.Sprintf(`{"action": "multiply", "x": %v, "y": %v, "answer", %v, "cached": false}`, x, y, answer)

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, response)
}

func extractParams(w http.ResponseWriter, r *http.Request) (x float64, y float64, err error) {

	x, err = strconv.ParseFloat(r.URL.Query().Get("x"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := fmt.Sprintf(`{"code": %v, "message": "%v"}`, 1, err.Error())
		fmt.Fprintf(w, response)
		return x, y, err
	}

	y, err = strconv.ParseFloat(r.URL.Query().Get("y"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := fmt.Sprintf(`{"code": %v, "message": "%v"}`, 1, err.Error())
		fmt.Fprintf(w, response)
		return x, y, err
	}

	return x, y, err
}
