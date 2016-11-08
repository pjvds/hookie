package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/handlers"
	"github.com/urfave/cli"
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

func main() {
	app := cli.NewApp()
	app.Name = "hookie"
	app.Usage = "webhook to script host"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "address",
			Value: "0.0.0.0:8080",
			Usage: "the address to listen on, for example: \"127.0.0.1:8080\"",
		},
	}
	app.Action = func(ctx *cli.Context) error {
		if len(ctx.Args()) == 0 {
			return fmt.Errorf("please specify the command to run, for example:\n\thookie my-script.sh")
		}

		address := ctx.String("address")
		command := ctx.Args()[0]
		args := ctx.Args()[1:]

		fmt.Printf("servning on %v\n", address)

		return http.ListenAndServe(address, handlers.LoggingHandler(os.Stdout, &CommandHandler{
			Command: command,
			Args:    args,
		}))
	}

	app.Run(os.Args)
}
