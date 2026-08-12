package commands

import "github.com/lanxre/kyokusu-cli/internal/models"

func HasCommand(cmds []models.Command, name string) bool {
	for _, cmd := range cmds {
		if cmd.Name == name {
			return true
		}
	}

	return false
}