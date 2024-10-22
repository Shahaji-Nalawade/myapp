package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLatencyLogger(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(50 * time.Millisecond) // Simulate some processing time
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	latencyLogger := LatencyLogger(handler)

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	latencyLogger.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Hello, World!")
}
