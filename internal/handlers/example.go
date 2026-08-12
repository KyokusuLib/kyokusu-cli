package handlers

import (
	"context"
	"fmt"

	"github.com/lanxre/kyokusu-cli/internal/models"
)

func RunWifi(
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