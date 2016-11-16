package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMethodFilterForwardsMatch(t *testing.T) {
	assert := assert.New(t)
	requests := []*http.Request{
		mustRawRequest("POST /url HTTP/1.1\n\n"),
	}

	for _, request := range requests {
		response := httptest.NewRecorder()

		invoked := false
		var upstreamRequest *http.Request
		var upstreamResponse http.ResponseWriter

		MethodsFilter{
			Methods: []string{"POST"},
			Handler: http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
				invoked = true
				upstreamRequest = request
				upstreamResponse = response
			}),
		}.ServeHTTP(response, request)

		assert.True(invoked, "did not invoke upstream hander for: "+request.URL.Path)
		assert.NotNil(upstreamRequest, "upstream handler did not receive request")
		assert.NotNil(upstreamRequest, "upstream handler did not receive response")
	}
}

func TestMethodsFilterDoesNotForwardsMismatches(t *testing.T) {
	assert := assert.New(t)
	requests := []*http.Request{
		mustRawRequest("get / HTTP/1.1\n\n"),
		mustRawRequest("POST / HTTP/1.1\n\n"),
		mustRawRequest("PUT /URL HTTP/1.1\n\n"),
		mustRawRequest("CONNECT /Url HTTP/1.1\n\n"),
	}

	for _, request := range requests {
		response := httptest.NewRecorder()

		invoked := false
		var upstreamRequest *http.Request
		var upstreamResponse http.ResponseWriter

		MethodsFilter{
			Methods: []string{"GET"},
			Handler: http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
				invoked = true
				upstreamRequest = request
				upstreamResponse = response
			}),
		}.ServeHTTP(response, request)

		assert.False(invoked, "did invoke upstream hander for: "+request.URL.String())
	}
}
