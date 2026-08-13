package internal

import (
	"context"
	"os"

	"github.com/lanxre/kyokusu-cli/internal/constants"
	"github.com/lanxre/kyokusu-cli/internal/handlers"
	"github.com/lanxre/kyokusu-cli/internal/models"
	"github.com/lanxre/kyokusu-cli/internal/utils"
)

var (
	globalOptions  []models.Definition
	globalCommands []models.Definition
)

func init() {
	globalOptions = []models.Definition{
		{
			Name:     "help",
			Short:    "Show help and exit",
			Terminal: true,
			Handler: func(ctx context.Context, in models.Input) error {
				return printHelp()
			},
		},
		{
			Name:     "version",
			Short:    "Show version and exit",
			Terminal: true,
			Handler:  handlers.RunVersion,
		},
	}

	globalCommands = []models.Definition{
		{
			Name:  "dmp",
			Short: "Dump database",
			Children: map[string]*models.Definition{
				"postgres": {
					Name:     "postgres",
					Short:    "Dump PostgreSQL database",
					Terminal: true,
					Handler:  handlers.RunDmp,
					Options: map[string]*models.Definition{
						"env": {
							Name:     "env",
							Short:    "Path to env file with database credentials",
							Terminal: true,
						},
						"o": {
							Name:     "o",
							Short:    "Output file path",
							Terminal: true,
						},
						"f": {
							Name:     "f",
							Short:    "Output format (sql, json)",
							Terminal: true,
						},
						"n": {
							Name:     "n",
							Short:    "Database name",
							Terminal: true,
						},
						"p": {
							Name:     "p",
							Short:    "Database port",
							Terminal: true,
						},
						"u": {
							Name:     "u",
							Short:    "User database",
							Terminal: true,
						},
						"psw": {
							Name:     "psw",
							Short:    "Password",
							Terminal: true,
						},
					},
				},
				"sqlite": {
					Name:     "sqlite",
					Short:    "Dump SQLite database",
					Terminal: true,
					Handler:  handlers.RunDmp,
					Options: map[string]*models.Definition{
						"env": {
							Name:     "env",
							Short:    "Path to env file with database credentials",
							Terminal: true,
						},
						"o": {
							Name:     "o",
							Short:    "Output file path",
							Terminal: true,
						},
						"f": {
							Name:     "f",
							Short:    "Output format (sql, json)",
							Terminal: true,
						},
						"db": {
							Name:     "db",
							Short:    "Path to SQLite database file",
							Terminal: true,
						},
					},
				},
			},
		},
	}
}

func printHelp() error {
	utils.PrintInit(os.Stdout, constants.RootMessage, globalOptions, globalCommands)
	return nil
}
