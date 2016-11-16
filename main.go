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
	app.Version = "v1.0"
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
		cli.BoolTFlag{
			Name:  "access-log",
			Usage: "print access log to stdout",
		},
		cli.StringFlag{
			Name:  "path",
			Value: "",
			Usage: "that pattern of the path to match",
		},
	}
	app.Action = func(ctx *cli.Context) error {
		if len(ctx.Args()) == 0 {
			return fmt.Errorf("please specify the command to run, for example:\n\thookie my-script.sh")
		}

		address := ctx.String("address")
		command := ctx.Args()[0]
		args := ctx.Args()[1:]

		fmt.Printf("serving on: %v\n", address)
		handler := http.Handler(&CommandHandler{
			Command: command,
			Args:    args,
		})

		if path := ctx.String("path"); len(path) > 0 {
			handler = &PathFilter{
				Path:    path,
				Handler: handler,
			}
		}

		if secret := ctx.String("github-secret"); len(secret) > 0 {
			fmt.Printf("github signature validation enabled\n")
			handler = &GithubSignatureValidator{
				Secret:  []byte(secret),
				Handler: handler,
			}
		}

		if log := ctx.BoolT("access-log"); log {
			handler = handlers.LoggingHandler(os.Stdout, handler)
		}

		if err := http.ListenAndServe(address, handler); err != nil {
			fmt.Printf("listen failure: %v\n", err)
			return err
		}

		return nil
	}

	app.Run(os.Args)
}
