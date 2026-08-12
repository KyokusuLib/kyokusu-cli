package internal

import (
	"context"
	"os"

	"github.com/lanxre/kyokusu-cli/internal/commands"
	"github.com/lanxre/kyokusu-cli/internal/constants"
	"github.com/lanxre/kyokusu-cli/internal/utils"
)

func Run(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	args := os.Args[1:]

	if len(args) == 0 {
		utils.PrintInit(os.Stdout, constants.RootMessage)
		return nil
	}

	input := commands.Parse(args)

	registry := NewRegistry()

	if err := registry.Resolve(input); err != nil {
		return err
	}

	return registry.Execute(ctx, input)
}
