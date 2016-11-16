package main

import "net/http"

type PathFilter struct {
	Path    string
	Handler http.Handler
}

func (this *PathFilter) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != this.Path {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	this.Handler.ServeHTTP(response, request)
}
