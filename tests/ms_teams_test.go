package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mw3tv123/go-notify/config"
	"github.com/mw3tv123/go-notify/server"
	"github.com/stretchr/testify/assert"
)

func TestMSTeamNotify(t *testing.T) {
	config.Init("development")
	router := server.NewRouter()

	var successBody = []byte(`{ "title": "test", "content": "test" }`)
	var invalidBody = []byte(`{ "invalid_field": "invalid)_value", "content": "test" }`)

	successRec := httptest.NewRecorder()
	failRec := httptest.NewRecorder()

	reqSuccess, _ := http.NewRequest("POST", "/api/ms_teams", bytes.NewBuffer(successBody))
	reqFail, _ := http.NewRequest("POST", "/api/ms_teams", bytes.NewBuffer(invalidBody))

	router.ServeHTTP(successRec, reqSuccess)
	router.ServeHTTP(failRec, reqFail)

	assert.Equal(t, 200, successRec.Code)
	assert.Equal(t, 406, failRec.Code)
}
