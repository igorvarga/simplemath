package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)
var basicmathtests = []struct {
	x        string
	y        string
	expected string
	handler  http.HandlerFunc
	name     string
	action   string
	path     string
}{
	{"2", "5", "7", AddHandler, "AddHandler", "add", "/add"},
	{"2", "5", "-3", SubtractHandler, "SubtractHandler", "subtract", "/subtract"},
	{"2", "5", "10", MultiplyHandler, "MultiplyHandler", "multiply", "/multiply"},
	{"10", "5", "2", DivideHandler, "DivideHandler", "divide", "/divide"},
}

func TestMathHandlers(t *testing.T) {

	for _, mt := range basicmathtests {
		t.Run(mt.name, func(t *testing.T) {
			values := url.Values{
				"x": {mt.x},
				"y": {mt.y}}

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

			expected := fmt.Sprintf(`{"action": "%v", "x": %v, "y": %v, "answer", %v, "cached": false}`, mt.action, mt.x, mt.y, mt.expected)
			if rr.Body.String() != expected {
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

	req, err := http.NewRequest("GET", url.String(), nil)
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

	req, err := http.NewRequest("GET", url.String(), nil)
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

func TestSubtractHandler(t *testing.T) {
	values := url.Values{
		"x": {"2"},
		"y": {"5"}}

	url := url.URL{Path: "/add", RawQuery: values.Encode()}

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SubtractHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"action": "subtract", "x": 2, "y": 5, "answer", -3, "cached": false}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

