package std

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
		{x: 2, y: 5, expected: 7, handler: AddHandler, name: "AddHandler", action: message.Add, path: "/add"},
		{x: 2, y: 5, expected: -3, handler: SubtractHandler, name: "SubtractHandler", action: message.Subtract, path: "/subtract"},
		{x: 2, y: 5, expected: 10, handler: MultiplyHandler, name: "MultiplyHandler", action: message.Multiply, path: "/multiply"},
		{x: 10, y: 5, expected: 2, handler: DivideHandler, name: "DivideHandler", action: message.Divide, path: "/divide"},
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

			expected, err := json.Marshal(message.ResultMessage{
				Action: mt.action,
				X:      mt.x,
				Y:      mt.y,
				Answer: mt.expected,
				Cached: false,
			})
			if err != nil {
				t.Errorf("Error encoding JSON")
			}

			es := string(expected)

			if rr.Body.String() != es {
				t.Errorf("Handler returned unexpected body: got %v want %v",
					rr.Body.String(), es)
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
