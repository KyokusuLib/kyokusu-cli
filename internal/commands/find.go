package commands

import "github.com/lanxre/kyokusu-cli/internal/models"

func FindOption(name string, options []models.Option) *models.Option {
	for _, option := range options {
		if option.Name == name {
			return &option
		}
	}
	return nil
}