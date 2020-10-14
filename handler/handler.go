package handler

import (
	"encoding/json"
	"fmt"
	"github.com/igorvarga/teletchcodechallenge/message"
	sm "github.com/igorvarga/teletchcodechallenge/simplemath"
	"net/http"
	"strconv"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {
	x, y, err := extractParams(w, r)
	if err != nil {
		return
	}

	answer := sm.Add(x, y)

	writeResult(w, x, y, answer, message.Add)
}

func SubtractHandler(w http.ResponseWriter, r *http.Request) {
	x, y, err := extractParams(w, r)
	if err != nil {
		return
	}

	answer := sm.Subtract(x, y)

	writeResult(w, x, y, answer, message.Subtract)
}

func DivideHandler(w http.ResponseWriter, r *http.Request) {
	x, y, err := extractParams(w, r)
	if err != nil {
		return
	}

	answer := sm.Divide(x, y)

	writeResult(w, x, y, answer, message.Divide)
}

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	x, y, err := extractParams(w, r)
	if err != nil {
		return
	}

	answer := sm.Multiply(x, y)

	writeResult(w, x, y, answer, message.Multiply)
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

func writeResult(w http.ResponseWriter, x float64, y float64, answer float64, action string) {
	result := message.ResultMessage{X: x, Y: y, Answer: answer}

	b, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}