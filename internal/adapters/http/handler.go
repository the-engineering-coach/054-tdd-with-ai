package http

import (
	"net/http"
)

type FlightHandler struct{}

func NewFlightHandler(repo interface{}) *FlightHandler {
	return &FlightHandler{}
}

func (h *FlightHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
}
