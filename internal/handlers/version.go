package handlers

import (
	"context"
	"fmt"

	"github.com/lanxre/kyokusu-cli/internal/constants"
	"github.com/lanxre/kyokusu-cli/internal/models"
)

func RunVersion(
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