package constants

import "github.com/lanxre/kyokusu-cli/internal/models"

const (
	ToolName = "KyokusuCli"
	Version  = "v0.0.1"
	Usage    = "kyokusu-cli"
)

type Dialect string

const (
	DialectSqlite Dialect = "sqlite"
	DialectPostgres Dialect = "postgres"
)


const (
	 BaseFormat = "sql"
	 BaseOutput = "backup"
)

var RootMessage = models.Root{
	Usage: Usage,
	Short: "Kyokusu CLI",
	Long: `KyokusuCli is a command-line tool manage 'KyokusuLib'.
It groups the project's operational commands under one entry point.`,
}

var (
	GlobalOptions  []models.Definition
	GlobalCommands []models.Definition
)