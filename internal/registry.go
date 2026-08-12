package internal

import (
	"context"
	"fmt"
	"os"

	"github.com/lanxre/kyokusu-cli/internal/commands"
	"github.com/lanxre/kyokusu-cli/internal/constants"
	"github.com/lanxre/kyokusu-cli/internal/models"
	"github.com/lanxre/kyokusu-cli/internal/utils"
)

func NewRegistry() *commands.Registry {
	registry := commands.NewRegistry()

	registry.RegisterGlobalOption(&commands.Definition{
		Name:     "help",
		Terminal: true,
		Handler:  runHelp,
	})

	registry.RegisterGlobalOption(&commands.Definition{
		Name:     "version",
		Terminal: true,
		Handler:  runVersion,
	})

	registry.RegisterCommand(&commands.Definition{
		Name: "dmp",
		Options: map[string]*commands.Definition{
			"u": {
				Name:     "u",
				Terminal: true,
			},
		},
		Handler:  runDmp,
	})

	registry.RegisterCommand(&commands.Definition{
		Name: "device",
		Children: map[string]*commands.Definition{
			"wifi": {
				Name: "wifi",
				Options: map[string]*commands.Definition{
					"rescan": {
						Name: "rescan",
					},
				},
				Handler: runWifi,
			},
		},
	})

	return registry
}

func runHelp(
	ctx context.Context,
	input models.Input,
) error {
	utils.PrintInit(os.Stdout, constants.RootMessage)
	return nil
}

func runVersion(
	ctx context.Context,
	input models.Input,
) error {
	fmt.Printf(
		"%s %s\n",
		constants.ToolName,
		constants.Version,
	)

	return nil
}

func runWifi(
	ctx context.Context,
	input models.Input,
) error {
	command := input.Command

	for command.Child != nil {
		command = command.Child
	}

	if command.Options == nil {
		fmt.Println("No options")
		return nil
	}

	for _, option := range command.Options {
		fmt.Printf("%s=%s\n", option.Name, option.Value)
	}

	return nil
}

func runDmp(
	ctx context.Context,
	input models.Input,
) error {
	fmt.Println("DUMP DATABASE")
	command := input.Command
	for _, option := range command.Options {
		fmt.Printf("%s=%s\n", option.Name, option.Value)
	}
	return nil
}
