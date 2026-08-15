package internal

import (
	"github.com/lanxre/kyokusu-cli/internal/commands"
	"github.com/lanxre/kyokusu-cli/internal/constants"
)

func NewRegistry() *commands.Registry {
	registry := commands.NewRegistry()

	for i := range constants.GlobalOptions {
		registry.RegisterGlobalOption(&constants.GlobalOptions[i])
	}

	for i := range constants.GlobalCommands {
		registry.RegisterCommand(&constants.GlobalCommands[i])
	}

	return registry
}
