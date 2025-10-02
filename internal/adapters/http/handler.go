package http

import (
	"encoding/json"
	"net/http"
)

type FlightHandler struct{}

func NewFlightHandler(repo interface{}) *FlightHandler {
	return &FlightHandler{}
}

func (h *FlightHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{"error": "origin parameter is required"})
}
