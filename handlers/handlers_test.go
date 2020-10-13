package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddHandler(t *testing.T) {
	u := fmt.Sprintf("/add?x=%v&=%v", 2, 5)

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"result": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
