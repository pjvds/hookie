package main

import "net/http"

type MethodsFilter struct {
	Methods []string
	Handler http.Handler
}

func (this *MethodsFilter) isAllowed(method string) bool {
	for _, allowed := range this.Methods {
		if method == allowed {
			return true
		}
	}
	return false
}

func (this *MethodsFilter) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if this.isAllowed(request.Method) {
		response.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	this.Handler.ServeHTTP(response, request)
}
