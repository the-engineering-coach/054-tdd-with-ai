package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func TestFlightHandler_NoOriginProvided(t *testing.T) {
	handler := NewFlightHandler(nil)

	req := httptest.NewRequest(http.MethodGet, "/flights", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var errResp ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&errResp)
	require.NoError(t, err)

	assert.Equal(t, "origin parameter is required", errResp.Error)
}
