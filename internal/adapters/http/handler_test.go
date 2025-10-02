package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlightHandler_NoOriginProvided(t *testing.T) {
	handler := NewFlightHandler(nil)

	req := httptest.NewRequest(http.MethodGet, "/flights", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
