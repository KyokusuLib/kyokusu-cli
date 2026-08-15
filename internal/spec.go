package internal

import (
	"github.com/lanxre/kyokusu-cli/internal/handlers"
)

func init() {
	RegisterGlobalOption("help", "Show help and exit", handlers.RunHelp)
	RegisterGlobalOption("version", "Show version and exit", handlers.RunVersion)

	RegisterCommand("dmp", "Dump database",
		WithSubcommand("postgres", "Dump PostgreSQL database",
			WithHandler(handlers.RunDmp),
			WithOption("env", "Path to env file with database credentials"),
			WithOption("o", "Output file path"),
			WithOption("f", "Output format (sql, json)"),
			WithOption("n", "Database name"),
			WithOption("p", "Database port"),
			WithOption("u", "User database"),
			WithOption("psw", "Password"),
		),
		WithSubcommand("sqlite", "Dump SQLite database",
			WithHandler(handlers.RunDmp),
			WithOption("env", "Path to env file with database credentials"),
			WithOption("o", "Output file path"),
			WithOption("f", "Output format (sql, json)"),
			WithOption("db", "Path to SQLite database file"),
		),
	)
	
	RegisterCommand("kill", "Kill running process",
		WithHandler(handlers.RunKill),
		WithOption("p", "Process port"),
	)
	
}
