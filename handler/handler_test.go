package handler

import (
	"encoding/json"
	"fmt"
	"github.com/igorvarga/teletchcodechallenge/message"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var basicmathtests = []struct {
	x        float64
	y        float64
	expected float64
	handler  http.HandlerFunc
	name     string
	action   string
	path     string
}{
	{2, 5, 7, AddHandler, "AddHandler", message.Add,"/add"},
	{2, 5, -3, SubtractHandler, "SubtractHandler", message.Subtract, "/subtract"},
	{2, 5, 10, MultiplyHandler, "MultiplyHandler", message.Multiply, "/multiply"},
	{10, 5, 2, DivideHandler, "DivideHandler", message.Divide, "/divide"},
}

func TestMathHandlers(t *testing.T) {

	for _, mt := range basicmathtests {
		t.Run(mt.name, func(t *testing.T) {
			values := url.Values{
				"x": {fmt.Sprintf("%v", mt.x)},
				"y": {fmt.Sprintf("%v", mt.y)}}

			url := url.URL{Path: mt.path, RawQuery: values.Encode()}

			req, err := http.NewRequest(http.MethodGet, url.String(), nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(mt.handler)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("Handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			// expected := fmt.Sprintf(`{"action": "%v", "x": %v, "y": %v, "answer", %v, "cached": false}`, mt.action, mt.x, mt.y, mt.expected)
			expected, err := json.Marshal(message.ResultMessage{
				X:      mt.x,
				Y:      mt.y,
				Answer: mt.expected,
				Cached: false,
			})
			if err != nil {
				t.Errorf("Error encoding JSON")
			}

			if rr.Body.String() != string(expected) {
				t.Errorf("Handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		})
	}
}

// TODO Refactor into table tests
func TestAddHandlerXMissing(t *testing.T) {
	values := url.Values{
		"y": {"5"}}

	url := url.URL{Path: "/add", RawQuery: values.Encode()}

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestAddHandlerYMissing(t *testing.T) {
	values := url.Values{
		"x": {"5"}}

	url := url.URL{Path: "/add", RawQuery: values.Encode()}

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}