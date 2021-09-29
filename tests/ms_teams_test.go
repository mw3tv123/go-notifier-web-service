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

func TestMSTeamsNotify(t *testing.T) {
	config.Init()
	router := server.NewRouter()
	successRec := httptest.NewRecorder()
	failRec := httptest.NewRecorder()

	var successBody = []byte(`{ "title": "test", "content": "test" }`)
	var invalidTitleBody = []byte(`{ "invalid_field": "invalid)_value", "content": "test" }`)
	var invalidChannelFieldBody = []byte(`{ "title": "test", "content": "test", "channels": "asdasd" }`)
	var invalidChannelBody = []byte(`{ "title": "test", "content": "test", "channels": ["tea"] }`)
	var titleLengthInvalidBody = []byte(`{ "title": "2c", "content": "test" }`)

	reqSuccess, _ := http.NewRequest("POST", "/message/ms_teams", bytes.NewBuffer(successBody))
	reqFail, _ := http.NewRequest("POST", "/message/ms_teams", bytes.NewBuffer(invalidTitleBody))
	router.ServeHTTP(successRec, reqSuccess)
	router.ServeHTTP(failRec, reqFail)
	assert.Equal(t, 200, successRec.Code)
	assert.Equal(t, 406, failRec.Code)

	reqFail, _ = http.NewRequest("POST", "/message/ms_teams", bytes.NewBuffer(invalidChannelFieldBody))
	router.ServeHTTP(failRec, reqFail)
	assert.Equal(t, 406, failRec.Code)

	reqFail, _ = http.NewRequest("POST", "/message/ms_teams", bytes.NewBuffer(invalidChannelBody))
	router.ServeHTTP(failRec, reqFail)
	assert.Equal(t, 406, failRec.Code)

	reqFail, _ = http.NewRequest("POST", "/message/ms_teams", bytes.NewBuffer(titleLengthInvalidBody))
	router.ServeHTTP(failRec, reqFail)
	assert.Equal(t, 406, failRec.Code)
}

func TestMSTeamsAlert(t *testing.T) {
	config.Init()
	router := server.NewRouter()

	var successBody = []byte(`{ "title": "test", "priority": 1, "service_name": "monitor a", "description": "Alert test a", "create_date": "2018-09-22T12:42:31+07:00" }`)
	var invalidBody = []byte(`{ "invalid_field": "invalid)_value", "content": "test" }`)
	var titleLengthInvalidBody = []byte(`{ "title": "2c", "content": "test" }`)

	successRec := httptest.NewRecorder()
	failRec := httptest.NewRecorder()
	failTitleRec := httptest.NewRecorder()

	reqSuccess, _ := http.NewRequest("POST", "/alert/ms_teams", bytes.NewBuffer(successBody))
	reqFail, _ := http.NewRequest("POST", "/alert/ms_teams", bytes.NewBuffer(invalidBody))
	reqTitleFail, _ := http.NewRequest("POST", "/alert/ms_teams", bytes.NewBuffer(titleLengthInvalidBody))

	router.ServeHTTP(successRec, reqSuccess)
	router.ServeHTTP(failRec, reqFail)
	router.ServeHTTP(failTitleRec, reqTitleFail)

	assert.Equal(t, 200, successRec.Code)
	assert.Equal(t, 406, failRec.Code)
	assert.Equal(t, 406, failTitleRec.Code)
}
