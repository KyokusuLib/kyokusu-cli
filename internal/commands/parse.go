package commands

import (
	"strings"

	"github.com/lanxre/kyokusu-cli/internal/models"
)

func Parse(args []string) models.Input {
	var input models.Input

	var current *models.Command

	for i := 0; i < len(args); i++ {
		arg := args[i]

		if strings.HasPrefix(arg, "-") {
			option := models.Option{
				Name: normalizeOption(strings.TrimLeft(arg, "-")),
			}

			if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
				option.Value = args[i+1]
				i++
			}

			if current == nil {
				input.Options = append(input.Options, option)
			} else {
				current.Options = append(current.Options, option)
			}

			continue
		}

		command := &models.Command{
			Name: arg,
		}

		if input.Command == nil {
			input.Command = command
			current = command
			continue
		}

		current.Child = command
		current = command
	}

	return input
}

func normalizeOption(name string) string {
	switch name {
	case "h":
		return "help"
	case "v":
		return "version"
	default:
		return name
	}
}