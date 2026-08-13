package constants

import "github.com/lanxre/kyokusu-cli/internal/models"

const (
	ToolName = "KyokusuCli"
	Version  = "v0.0.1"
	Usage    = "kyokusu-cli"
)

var RootMessage = models.Root{
	Usage: Usage,
	Short: "Kyokusu CLI",
	Long: `KyokusuCli is a command-line tool manage 'KyokusuLib'.
It groups the project's operational commands under one entry point.`,
}