package server

import (
	"encoding/json"
	"fmt"
	"github.com/igorvarga/teltechcodechallenge/cache"
	"github.com/igorvarga/teltechcodechallenge/message"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func TestMathHTTPHandlers(t *testing.T) {

	c := cache.NewCache(time.Minute, 5 * time.Minute)
	s := NewSimpleMathServer(c)

	basicmathtests := []struct {
		x        float64
		y        float64
		expected float64
		handler  http.HandlerFunc
		name     string
		action   string
		path     string
	}{
		{x: 2, y: 5, expected: 7, handler: s.AddHandler, name: "AddHandler", action: message.ActionAdd, path: "/add"},
		{x: 2, y: 5, expected: -3, handler: s.SubtractHandler, name: "SubtractHandler", action: message.ActionSubtract, path: "/subtract"},
		{x: 2, y: 5, expected: 10, handler: s.MultiplyHandler, name: "MultiplyHandler", action: message.ActionMultiply, path: "/multiply"},
		{x: 10, y: 5, expected: 2, handler: s.DivideHandler, name: "DivideHandler", action: message.ActionDivide, path: "/divide"},
	}

	for _, mt := range basicmathtests {
		t.Run(mt.name, func(t *testing.T) {
			values := url.Values{
				"x": {fmt.Sprintf("%v", mt.x)},
				"y": {fmt.Sprintf("%v", mt.y)}}

			newurl := url.URL{Path: mt.path, RawQuery: values.Encode()}

			req, err := http.NewRequest(http.MethodGet, newurl.String(), nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := mt.handler
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
	s := NewSimpleMathServer(nil)

	values := url.Values{
		"y": {"5"}}

	newurl := url.URL{Path: "/add", RawQuery: values.Encode()}

	req, err := http.NewRequest(http.MethodGet, newurl.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.AddHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestAddHandlerYMissing(t *testing.T) {
	s := NewSimpleMathServer(nil)

	values := url.Values{
		"x": {"5"}}

	newurl := url.URL{Path: "/add", RawQuery: values.Encode()}

	req, err := http.NewRequest(http.MethodGet, newurl.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.AddHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestCacheMiddleware(t *testing.T) {
	c := cache.NewCache(time.Minute, 5 * time.Minute)
	s := NewSimpleMathServer(c)

	values := url.Values{
		"x": {"5"},
		"y": {"5"}}

	newurl := url.URL{Path: "/add", RawQuery: values.Encode()}

	req, err := http.NewRequest(http.MethodGet, newurl.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := s.CacheMiddleware(s.AddHandler)
	handler.ServeHTTP(rr, req)

	// First call test
	expected, err := json.Marshal(message.ResultMessage{
		Action: message.ActionAdd,
		X:      5,
		Y:      5,
		Answer: 10,
	})
	if err != nil {
		t.Errorf("Error encoding JSON")
	}

	es := string(expected)

	if rr.Body.String() != es {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), es)
	}

	// test cached result
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	expected, err = json.Marshal(message.ResultMessage{
		Action: message.ActionAdd,
		X:      5,
		Y:      5,
		Answer: 10,
		Cached: "true",
	})
	if err != nil {
		t.Errorf("Error encoding JSON")
	}

	es = string(expected)

	if rr.Body.String() != es {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), es)
	}

}

func TestCacheMiddlewareExpired(t *testing.T) {
	c := cache.NewCache(100 * time.Millisecond, 5 * time.Minute)
	s := NewSimpleMathServer(c)

	values := url.Values{
		"x": {"5"},
		"y": {"5"}}

	newurl := url.URL{Path: "/add", RawQuery: values.Encode()}

	req, err := http.NewRequest(http.MethodGet, newurl.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := s.CacheMiddleware(s.AddHandler)
	handler.ServeHTTP(rr, req)

	// First call test
	expected, err := json.Marshal(message.ResultMessage{
		Action: message.ActionAdd,
		X:      5,
		Y:      5,
		Answer: 10,
	})
	if err != nil {
		t.Errorf("Error encoding JSON")
	}

	es := string(expected)

	if rr.Body.String() != es {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), es)
	}

	time.Sleep(200 * time.Millisecond)
	
	// test cached result
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	expected, err = json.Marshal(message.ResultMessage{
		Action: message.ActionAdd,
		X:      5,
		Y:      5,
		Answer: 10,
		Cached: "true",
	})
	if err != nil {
		t.Errorf("Error encoding JSON")
	}

	es = string(expected)

	if rr.Body.String() != es {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), es)
	}

}