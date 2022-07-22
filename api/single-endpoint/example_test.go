package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloRoute(t *testing.T) {
	var helloResponse HelloResponse

	engine := SetupEngine()

	request, _ := http.NewRequest("GET", "/hello", nil)
	response := httptest.NewRecorder()

	engine.GET("/hello", HelloHandler)
	engine.ServeHTTP(response, request)

	json.Unmarshal(response.Body.Bytes(), &helloResponse)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "world", helloResponse.Message)
}

func TestHelloRouteFail(t *testing.T) {
	var helloResponse HelloResponse

	engine := SetupEngine()

	request, _ := http.NewRequest("GET", "/hello", nil)
	response := httptest.NewRecorder()

	engine.GET("/hello", HelloHandler)
	engine.ServeHTTP(response, request)

	json.Unmarshal(response.Body.Bytes(), &helloResponse)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.NotEqual(t, "code", helloResponse.Message)
}
