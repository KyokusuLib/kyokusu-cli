package internal

import (
	"github.com/lanxre/kyokusu-cli/internal/commands"
	"github.com/lanxre/kyokusu-cli/internal/constants"
	"github.com/lanxre/kyokusu-cli/internal/models"
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

func RegisterGlobalOption(name, description string, fn models.Handler) {
	def := models.Definition{
		Name:     name,
		Short:    description,
		Terminal: true,
		Handler:  fn,
	}
	constants.GlobalOptions = append(constants.GlobalOptions, def)
}

type CommandOption func(def *models.Definition)

func RegisterCommand(name, description string, options ...CommandOption) {
	def := &models.Definition{
		Name:  name,
		Short: description,
	}

	for _, option := range options {
		option(def)
	}

	constants.GlobalCommands = append(constants.GlobalCommands, *def)
}

func WithHandler(fn models.Handler) CommandOption {
	return func(def *models.Definition) {
		def.Handler = fn
		def.Terminal = true
	}
}

func WithOption(name, description string) CommandOption {
	return func(def *models.Definition) {
		if def.Options == nil {
			def.Options = make(map[string]*models.Definition)
		}

		def.Options[name] = &models.Definition{
			Name:     name,
			Short:    description,
			Terminal: true,
		}
	}
}

func WithSubcommand(name, description string, options ...CommandOption) CommandOption {
	return func(def *models.Definition) {
		if def.Children == nil {
			def.Children = make(map[string]*models.Definition)
		}

		child := &models.Definition{
			Name:  name,
			Short: description,
		}

		for _, option := range options {
			option(child)
		}

		def.Children[name] = child
	}
}
