package internal

import (
	"github.com/lanxre/kyokusu-cli/internal/commands"
)

func NewRegistry() *commands.Registry {
	registry := commands.NewRegistry()

	for i := range globalOptions {
		registry.RegisterGlobalOption(&globalOptions[i])
	}

	for i := range globalCommands {
		registry.RegisterCommand(&globalCommands[i])
	}

	return registry
}
