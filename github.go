package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type GithubSecretValidator struct {
	Handler http.Handler
	Secret  []byte
}

func verifySignature(secret []byte, signature string, body []byte) bool {
	// example: sha1=21178161e2fc27f5598342b6759e6a395a1fd008
	const prefix = "sha1="
	const lenght = 45

	if len(signature) != lenght {
		fmt.Fprintf(os.Stderr, "signature does not have expected lenght of %v characters: %v\n", lenght, signature)
		return false
	}

	if !strings.HasPrefix(signature, prefix) {
		fmt.Fprintf(os.Stderr, "signature does not have expected prefix [sha1=]: %v\n", signature)
		return false
	}

	actual, err := hex.DecodeString(signature[len(prefix):])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error decoding signature: %v\n", err)
		return false
	}

	return hmac.Equal(sign(secret, body), actual)
}

func sign(secret, data []byte) []byte {
	hash := hmac.New(sha1.New, secret)
	hash.Write(data)
	return hash.Sum(nil)
}

func (this *GithubSecretValidator) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	signature := request.Header.Get("X-Hub-Signature")
	if len(signature) == 0 {
		fmt.Fprint(os.Stderr, "missing signature\n")
		http.Error(response, "forbidden", http.StatusForbidden)
		return
	}

	buffer := new(bytes.Buffer)
	if _, err := io.Copy(buffer, request.Body); err != nil {
		fmt.Fprintf(os.Stderr, "request body read error: %v\n", err)
		http.Error(response, "internal server error", 500)
		return
	}
	request.Body = ioutil.NopCloser(buffer)

	if !verifySignature(this.Secret, signature, buffer.Bytes()) {
		fmt.Fprint(os.Stderr, "invalid signature\n")
		http.Error(response, "forbidden", http.StatusForbidden)
		return
	}

	this.Handler.ServeHTTP(response, request)
}
