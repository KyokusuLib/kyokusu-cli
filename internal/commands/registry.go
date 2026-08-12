package commands

import (
	"context"
	"fmt"

	"github.com/lanxre/kyokusu-cli/internal/models"
)

type Handler func(
	ctx context.Context,
	input models.Input,
) error

type Definition struct {
	Name     string
	Terminal bool
	Handler  Handler
	Children map[string]*Definition
	Options  map[string]*Definition
}

type Registry struct {
	commands      map[string]*Definition
	globalOptions map[string]*Definition
}

func NewRegistry() *Registry {
	return &Registry{
		commands:      make(map[string]*Definition),
		globalOptions: make(map[string]*Definition),
	}
}

func (r *Registry) RegisterCommand(definition *Definition) {
	r.commands[definition.Name] = definition
}

func (r *Registry) RegisterGlobalOption(definition *Definition) {
	r.globalOptions[definition.Name] = definition
}

func (r *Registry) Resolve(input models.Input) error {
	for _, option := range input.Options {
		if _, ok := r.globalOptions[option.Name]; !ok {
			return fmt.Errorf("unknown option: %s", option.Name)
		}
	}

	if input.Command == nil {
		return nil
	}

	definition, ok := r.commands[input.Command.Name]
	if !ok {
		return fmt.Errorf("unknown command: %s", input.Command.Name)
	}

	return resolveCommand(input.Command, definition)
}

func resolveCommand(
	command *models.Command,
	definition *Definition,
) error {
	for _, option := range command.Options {
		if _, ok := definition.Options[option.Name]; !ok {
			return fmt.Errorf(
				"unknown option '%s' for command '%s'",
				option.Name,
				command.Name,
			)
		}
	}

	if command.Child == nil {
		return nil
	}

	child, ok := definition.Children[command.Child.Name]
	if !ok {
		return fmt.Errorf(
			"unknown subcommand '%s' for command '%s'",
			command.Child.Name,
			command.Name,
		)
	}

	return resolveCommand(command.Child, child)
}

func (r *Registry) Execute(
	ctx context.Context,
	input models.Input,
) error {
	for _, option := range input.Options {
		definition := r.globalOptions[option.Name]

		if err := definition.Handler(ctx, input); err != nil {
			return err
		}

		if definition.Terminal {
			return nil
		}
	}

	if input.Command == nil {
		return nil
	}

	definition := r.commands[input.Command.Name]

	return executeCommand(ctx, input, input.Command, definition)
}

func executeCommand(
	ctx context.Context,
	input models.Input,
	command *models.Command,
	definition *Definition,
) error {
	if command.Child != nil {
		child := definition.Children[command.Child.Name]

		return executeCommand(
			ctx,
			input,
			command.Child,
			child,
		)
	}

	if definition.Handler == nil {
		return fmt.Errorf(
			"command '%s' has no handler",
			command.Name,
		)
	}

	return definition.Handler(ctx, input)
}