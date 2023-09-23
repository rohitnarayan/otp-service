package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/rohitnarayan/otp-service/internal/config"
	"github.com/rohitnarayan/otp-service/internal/postgres"
	"github.com/rohitnarayan/otp-service/internal/server"
)

func main() {
	config.Init()

	app := cli.NewApp()
	app.Name = "OTP Service"

	migrationConfig := postgres.MigrationConfig{
		Driver: "postgres",
		URL:    postgres.ConnectionURL(config.App.Database.Postgres),
		DBName: config.App.Database.Postgres.DatabaseName,
		Path:   "file://./migrations",
	}

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
		{
			Name:  "create-migration",
			Usage: "create migration files",
			Action: func(c *cli.Context) error {
				return postgres.CreateMigration(c.Args().Get(0), &migrationConfig)
			},
		},
		{
			Name:  "migrate",
			Usage: "run db migrations",
			Action: func(c *cli.Context) error {
				return postgres.RunDatabaseMigrations(&migrationConfig)
			},
		},
		{
			Name:  "rollback",
			Usage: "rollback db migrations",
			Action: func(c *cli.Context) error {
				return postgres.RollbackLatestMigration(&migrationConfig)
			},
		},
		{
			Name:    "drop-db",
			Aliases: []string{"m"},
			Usage:   "drop database",
			Action: func(c *cli.Context) error {
				return postgres.DropDatabase(&migrationConfig)
			},
		},
		{
			Name:  "create-db",
			Usage: "create database",
			Action: func(c *cli.Context) error {
				return postgres.CreateDatabase(&migrationConfig)
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
