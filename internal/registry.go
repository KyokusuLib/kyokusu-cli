package internal

import (
	"github.com/lanxre/kyokusu-cli/internal/commands"
	"github.com/lanxre/kyokusu-cli/internal/handlers"
)

func NewRegistry() *commands.Registry {
	registry := commands.NewRegistry()

	registry.RegisterGlobalOption(&commands.Definition{
		Name:     "help",
		Terminal: true,
		Handler:  handlers.RunHelp,
	})

	registry.RegisterGlobalOption(&commands.Definition{
		Name:     "version",
		Terminal: true,
		Handler:  handlers.RunVersion,
	})

	registry.RegisterCommand(&commands.Definition{
		Name: "dmp",
		Options: map[string]*commands.Definition{
			"u": {
				Name:     "u",
				Terminal: true,
			},
		},
		Handler:  handlers.RunDmp,
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
				Handler: handlers.RunWifi,
			},
		},
	})

	return registry
}

