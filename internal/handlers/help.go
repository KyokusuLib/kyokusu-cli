package handlers

import (
	"context"
	"os"

	"github.com/lanxre/kyokusu-cli/internal/constants"
	"github.com/lanxre/kyokusu-cli/internal/models"
	"github.com/lanxre/kyokusu-cli/internal/utils"
	
)

func RunHelp(
	ctx context.Context,
	input models.Input,
) error {
	utils.PrintInit(os.Stdout, constants.RootMessage, constants.GlobalOptions, constants.GlobalCommands)
	return nil
}