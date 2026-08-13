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
			Options: map[string]*models.Definition{
				"p": {
					Name:     "p",
					Short:    "Path to env file with database credentials",
					Terminal: true,
				},
			},
			Handler: handlers.RunDmp,
		},
	}
}

func printHelp() error {
	utils.PrintInit(os.Stdout, constants.RootMessage, globalOptions, globalCommands)
	return nil
}
