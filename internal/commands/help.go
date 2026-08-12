package commands

import "github.com/lanxre/kyokusu-cli/internal/models"

func IsCommandHelp(cmds []models.Command) bool {
	if len(cmds) == 0 {
		return false
	}

	return HasCommand(cmds, "help")
}