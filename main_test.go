package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistanceHandlerNotPost(t *testing.T) {

	req, err := http.NewRequest("GET", "/api/v1/distance", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(distanceHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, `{"Success":false,"Message":"Only HTTP POST is supported"}`, rr.Body.String())

}

func TestDistanceHandlerPostNilBodyError(t *testing.T) {

	req, err := http.NewRequest("POST", "/api/v1/distance", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(distanceHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, `{"Success":false,"Message":"HTTP POST body must contain a JSON object with keys 'source' and 'target'"}`, rr.Body.String())

}

func TestDistanceHandlerInvalidJSON(t *testing.T) {

	req, err := http.NewRequest("POST", "/api/v1/distance", strings.NewReader(`bad json`))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(distanceHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, `{"Success":false,"Message":"Error decoding JSON http post, expected valid JSON body: invalid character 'b' looking for beginning of value"}`, rr.Body.String())

}

func TestDistanceHandlerMissingSourceKeyPost(t *testing.T) {

	req, err := http.NewRequest("POST", "/api/v1/distance", strings.NewReader(`{"bad": "POST"}`))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(distanceHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, `{"Success":false,"Message":"HTTP POST must have JSON object with key called 'source' containing a comma delimited list of source strings to be mapped"}`, rr.Body.String())

}

func TestDistanceHandlerMissingTargetKeyPost(t *testing.T) {

	json := `{"source": "some data, some more data"}`
	reader := strings.NewReader(json)

	req, err := http.NewRequest("POST", "/api/v1/distance", reader)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(distanceHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, `{"Success":false,"Message":"HTTP POST must have JSON object with key called 'target' containing a comma delimited list of target strings to be mapped to"}`, rr.Body.String())

}
