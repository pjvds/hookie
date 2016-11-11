package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/urfave/cli"
)

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
		cli.StringFlag{
			Name:  "github-secret",
			Usage: "validates the incoming request with the X-Hub-Signature header value set by Github",
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
    handler := &CommandHandler{
			Command: command,
			Args:    args,
		}

    if secret := ctx.String("github-secret"); len(secret) > 0{
      fmt.Printf("github signature validation enabled\n")
      handler := &GithubSecretValidator{
        Secret: []byte(secret),
        Handler: handler,
      }
    }

		return http.ListenAndServe(address, handlers.LoggingHandler(os.Stdout, handlers.LoggingHandler(os.Stdout, handler))
	}

	app.Run(os.Args)
}
