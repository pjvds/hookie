package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

type CommandHandler struct {
	Command string
	Args    []string
}

func (this *CommandHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	cmd := exec.Command(this.Command, this.Args...)

	cmd.Stdout = io.MultiWriter(os.Stdout, response)
	cmd.Stderr = io.MultiWriter(os.Stderr, response)
	cmd.Stdin = request.Body

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		response.WriteHeader(500)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
	}
}
