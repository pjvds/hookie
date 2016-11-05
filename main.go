package main

import (
	"fmt"
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
	if err := cmd.Run(); err != nil {
		http.Error(response, err.Error(), 500)
		return
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "hookie"
	app.Usage = "webhook to script host"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "address",
			Value: ":8080",
			Usage: "the http addres to bind to, for example: \"127.0.0.1:8000\"",
		},
	}
	app.Action = func(ctx *cli.Context) error {
		if len(ctx.Args()) == 0 {
			return fmt.Errorf("missing command")
		}

		address := ctx.String("address")
		command := ctx.Args()[0]
		args := ctx.Args()[1:]

		fmt.Printf("listening on %v\n", address)

		return http.ListenAndServe(address, handlers.LoggingHandler(os.Stdout, &CommandHandler{
			Command: command,
			Args:    args,
		}))
	}

	app.Run(os.Args)
}
