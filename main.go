package main

import (
	"github.com/rohitnarayan/otp-service/internal/server"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "OTP Service"

	app.Commands = []cli.Command{
		{
			Name:    "server",
			Usage:   "start the OTP service",
			Aliases: []string{"serve", "server", "start"},
			Action: func(ctx *cli.Context) error {
				server.Server()
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("unable to start the server")
	}
}
