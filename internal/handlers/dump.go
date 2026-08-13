package handlers

import (
	"context"
	"fmt"

	"github.com/lanxre/kyokusu-cli/internal/models"
)


func RunDmp(
	ctx context.Context,
	input models.Input,
) error {
	fmt.Println("[DUMP DATABASE]")

	command := input.Command

	for _, option := range command.Options {
		fmt.Printf("%s=%s\n", option.Name, option.Value)
	}
	
	return nil
}
