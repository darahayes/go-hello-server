package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/hello", nil)
	response := httptest.NewRecorder()

	router := newRouter()
	router.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
