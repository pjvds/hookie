package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathFilterForwardsMatch(t *testing.T) {
	assert := assert.New(t)
	requests := []*http.Request{
		mustRawRequest("POST /url HTTP/1.1\n\n"),
	}

	for _, request := range requests {
		invoked := false
		var upstreamRequest *http.Request
		var upstreamResponse http.ResponseWriter

		pathFilter := PathFilter{
			Path: "/url",
			Handler: http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
				invoked = true
				upstreamRequest = request
				upstreamResponse = response
			}),
		}

		response := httptest.NewRecorder()
		pathFilter.ServeHTTP(response, request)

		assert.True(invoked, "did not invoke upstream hander for: "+request.URL.Path)
		assert.NotNil(upstreamRequest, "upstream handler did not receive request")
		assert.NotNil(upstreamRequest, "upstream handler did not receive response")
	}
}

func TestPathFilterDoesNotForwardsMismatches(t *testing.T) {
	assert := assert.New(t)
	requests := []*http.Request{
		mustRawRequest("POST / HTTP/1.1\n\n"),
		mustRawRequest("POST /Url HTTP/1.1\n\n"),
		mustRawRequest("POST /URL HTTP/1.1\n\n"),
	}

	for _, request := range requests {
		invoked := false
		var upstreamRequest *http.Request
		var upstreamResponse http.ResponseWriter

		pathFilter := PathFilter{
			Path: "/url",
			Handler: http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
				invoked = true
				upstreamRequest = request
				upstreamResponse = response
			}),
		}

		response := httptest.NewRecorder()
		pathFilter.ServeHTTP(response, request)

		assert.False(invoked, "did invoke upstream hander for: "+request.URL.String())
	}
}
