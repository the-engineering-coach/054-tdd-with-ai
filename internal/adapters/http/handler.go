package http

import (
	"encoding/json"
	"net/http"

	"the-engineering-coach/tdd-with-ai/internal/ports"
)

type FlightHandler struct {
	service ports.FlightService
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewFlightHandler(service ports.FlightService) *FlightHandler {
	return &FlightHandler{service: service}
}

func (h *FlightHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	origin := r.URL.Query().Get("origin")
	if origin == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "origin parameter is required"})
		return
	}

	flights, err := h.service.SearchByOrigin(r.Context(), origin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(flights)
}
