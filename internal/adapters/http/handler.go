package http

import (
	"encoding/json"
	"net/http"

	"the-engineering-coach/tdd-with-ai/internal/ports"
)

type FlightHandler struct {
	service ports.FlightService
}

func NewFlightHandler(service ports.FlightService) *FlightHandler {
	return &FlightHandler{service: service}
}

func (h *FlightHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	origin := r.URL.Query().Get("origin")
	if origin == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "origin parameter is required"})
		return
	}

	_, err := h.service.SearchByOrigin(r.Context(), origin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
}
