package server

import (
	"encoding/json"
	"fmt"
	. "github.com/igorvarga/simplemath/cache"
	"github.com/igorvarga/simplemath/message"
	"github.com/igorvarga/simplemath/simplemath"
	"log"
	"net/http"
	"strconv"
)

type SimpleMathServer struct {
	cache Cache
}

func NewSimpleMathServer(c Cache) *SimpleMathServer {
	return &SimpleMathServer{cache: c}
}

func (s *SimpleMathServer) AddHandler(w http.ResponseWriter, r *http.Request) {
	x, y, err := s.extractParams(w, r)
	if err != nil {
		return
	}

	answer := simplemath.Add(x, y)

	result := message.ResultMessage{
		Action: message.ActionAdd,
		X:      x,
		Y:      y,
		Answer: answer,
	}

	s.writeJSON(w, http.StatusOK, result)

	s.cacheResult(x, y, result)
}

func (s *SimpleMathServer) SubtractHandler(w http.ResponseWriter, r *http.Request) {
	x, y, err := s.extractParams(w, r)
	if err != nil {
		return
	}

	answer := simplemath.Subtract(x, y)

	result := message.ResultMessage{
		Action: message.ActionSubtract,
		X:      x,
		Y:      y,
		Answer: answer,
	}

	s.writeJSON(w, http.StatusOK, result)

	s.cacheResult(x, y, result)
}

func (s *SimpleMathServer) DivideHandler(w http.ResponseWriter, r *http.Request) {
	x, y, err := s.extractParams(w, r)
	if err != nil {
		return
	}

	answer := simplemath.Divide(x, y)

	result := message.ResultMessage{
		Action: message.ActionDivide,
		X:      x,
		Y:      y,
		Answer: answer,
	}

	s.writeJSON(w, http.StatusOK, result)

	s.cacheResult(x, y, result)
}

func (s *SimpleMathServer) MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	x, y, err := s.extractParams(w, r)
	if err != nil {
		return
	}

	answer := simplemath.Multiply(x, y)

	result := message.ResultMessage{
		Action: message.ActionMultiply,
		X:      x,
		Y:      y,
		Answer: answer,
	}

	s.writeJSON(w, http.StatusOK, result)

	s.cacheResult(x, y, result)
}

func (s *SimpleMathServer) cacheResult(x float64, y float64, result message.ResultMessage) {
	key := fmt.Sprint(x, ":", y)

	_, ok := s.cache.Load(key)
	if ok {
		return
	}

	result.Cached = "true"

	b, err := json.Marshal(result)
	if err != nil {
		log.Printf("Marshalling result error: %v", err.Error())
		return
	}

	s.cache.Store(key, b)
}

func (s *SimpleMathServer) CacheMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		x, y, err := s.extractParams(w, r)
		if err != nil {
			log.Printf("Unable to extract params for query %v", r.URL.RawQuery)
			return
		}

		key := fmt.Sprint(x, ":", y)

		if item, ok := s.cache.Load(key); ok {
			b := item.Value().([]byte)
			log.Println("Serving response from the cache.")

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			_, err = w.Write(b)
			if err != nil {
				log.Fatal("Unable to write JSON response.")
			}
			return
		}

		log.Println("Cache item not found, calling next handler.")

		next.ServeHTTP(w, r)

	}
}

func (s *SimpleMathServer) extractParams(w http.ResponseWriter, r *http.Request) (x float64, y float64, err error) {

	x, err = strconv.ParseFloat(r.URL.Query().Get("x"), 64)
	if err != nil {
		result := message.ErrorMessage{
			Code:  "1",
			Error: "Error parsing x parameter",
		}

		log.Printf("Error parsing x parameter from query: %v", r.URL.RawQuery)

		s.writeJSON(w, http.StatusBadRequest, result)

		return x, y, err
	}

	y, err = strconv.ParseFloat(r.URL.Query().Get("y"), 64)
	if err != nil {
		result := message.ErrorMessage{
			Code:  "1",
			Error: "Error parsing y parameter",
		}

		log.Printf("Error parsing y parameter from query: %v", r.URL.RawQuery)

		s.writeJSON(w, http.StatusBadRequest, result)

		return x, y, err
	}

	return x, y, err
}

func (s *SimpleMathServer) writeJSON(w http.ResponseWriter, status int, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_, err = w.Write(b)
	if err != nil {
		log.Fatal("Unable to write JSON response.")
	}

}
