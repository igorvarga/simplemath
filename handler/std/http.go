package std

import (
	"encoding/json"
	"github.com/igorvarga/teletchcodechallenge/message"
	"github.com/igorvarga/teletchcodechallenge/simplemath"
	"net/http"
	"strconv"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {
	x, y, err := extractParams(w, r)
	if err != nil {
		return
	}

	answer := simplemath.Add(x, y)

	result := message.ResultMessage{
		Action: message.ActionAdd,
		X:      x,
		Y:      y,
		Answer: answer,
		Cached: false,
	}

	writeJSON(w, http.StatusOK, result)
}

func SubtractHandler(w http.ResponseWriter, r *http.Request) {
	x, y, err := extractParams(w, r)
	if err != nil {
		return
	}

	answer := simplemath.Subtract(x, y)

	result := message.ResultMessage{
		Action: message.ActionSubtract,
		X:      x,
		Y:      y,
		Answer: answer,
		Cached: false,
	}

	writeJSON(w, http.StatusOK, result)
}

func DivideHandler(w http.ResponseWriter, r *http.Request) {
	x, y, err := extractParams(w, r)
	if err != nil {
		return
	}

	answer := simplemath.Divide(x, y)

	result := message.ResultMessage{
		Action: message.ActionDivide,
		X:      x,
		Y:      y,
		Answer: answer,
		Cached: false,
	}

	writeJSON(w, http.StatusOK, result)
}

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	x, y, err := extractParams(w, r)
	if err != nil {
		return
	}

	answer := simplemath.Multiply(x, y)

	result := message.ResultMessage{
		Action: message.ActionMultiply,
		X:      x,
		Y:      y,
		Answer: answer,
		Cached: false,
	}

	writeJSON(w, http.StatusOK, result)
}

func extractParams(w http.ResponseWriter, r *http.Request) (x float64, y float64, err error) {

	x, err = strconv.ParseFloat(r.URL.Query().Get("x"), 64)
	if err != nil {
		result := message.ErrorMessage{
			Code:  "1",
			Error: "Error parsing x parameter",
		}

		writeJSON(w, http.StatusBadRequest, result)

		return x, y, err
	}

	y, err = strconv.ParseFloat(r.URL.Query().Get("y"), 64)
	if err != nil {
		result := message.ErrorMessage{
			Code:  "1",
			Error: "Error parsing y parameter",
		}

		writeJSON(w, http.StatusBadRequest, result)

		return x, y, err
	}

	return x, y, err
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(b)
}
