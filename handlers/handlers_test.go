package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestAddHandler(t *testing.T) {
	values := url.Values{
		"x": {"2"},
		"y": {"5"}}

	url := url.URL{Path: "/add", RawQuery: values.Encode()}

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"action": "add", "x": 2, "y": 5, "answer", 7, "cached": false}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
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