package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

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

		return http.ListenAndServe(address, http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
			cmd := exec.Command(command, args...)
			if err := cmd.Run(); err != nil {
				fmt.Printf("command execution failed: %v", err.Error())
				http.Error(response, err.Error(), 500)
				return
			}
		}))
	}

	app.Run(os.Args)
}
