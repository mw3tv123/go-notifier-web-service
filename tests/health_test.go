package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mw3tv123/go-notify/server"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	router := server.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
