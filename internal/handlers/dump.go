package handlers

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/lanxre/kyokusu-cli/internal/commands"
	"github.com/lanxre/kyokusu-cli/internal/constants"
	"github.com/lanxre/kyokusu-cli/internal/models"
	"github.com/lanxre/kyokusu-cli/internal/utils"
)


func RunDmp(
	ctx context.Context,
	input models.Input,
) error {
	fmt.Print("[DUMP DATABASE]:")

	command := input.Command

	dialectDb := command.Child.Name

	switch dialectDb {
		case string(constants.DialectSqlite):
			if err := DumpSqlite(ctx, command.Child.Options); err != nil {
				return err
			}
		case string(constants.DialectPostgres):
			if err := DumpPostgres(ctx, command.Child.Options); err != nil {
				return err
			}
		default:
			return nil
	}

	fmt.Print(" Done\n")
	return nil
}


func DumpPostgres(ctx context.Context, options []models.Option) error {

	pgDumpPath, err := exec.LookPath("pg_dump")

	if err != nil {
		return fmt.Errorf("pg_dump not found: %w", err)
	}

	env := commands.FindOption("env", options)

	out, err := getOutput(options)
	if err != nil {
		return err
	}

	if env != nil {
		config, err := utils.ParsePostgresConfig(env.Value)

		if err != nil {
			return err
		}

		if err := utils.DumpPostgresConfig(pgDumpPath, config, out); err != nil {
			return err
		}

		return nil
		
	} else {

		port := commands.FindOption("p", options)
		user := commands.FindOption("u", options)
		name := commands.FindOption("n", options)
		password := commands.FindOption("psw", options)
		
		config := utils.PostgresConfig{
			Host:         "postgres",
			Port:         "5432",
			User:         "postgres",
			Password:     "postgres",
			DatabaseName: "postgres",
		}

		if port != nil {
			config.Port = port.Value
		}

		if user != nil {
			config.User = user.Value
		}

		if name != nil {
			config.DatabaseName = name.Value
		}

		if password != nil {
			config.Password = password.Value
		}

		if err := utils.DumpPostgresConfig(pgDumpPath, config, out); err != nil {
			return err
		}

		return nil
	}
}

func DumpSqlite(ctx context.Context, options []models.Option) error {
	sqlite3Path, err := exec.LookPath("sqlite3")
	if err != nil {
		return fmt.Errorf("sqlite3 not found: %w", err)
	}

	out, err := getOutput(options)
	if err != nil {
		return err
	}

	env := commands.FindOption("env", options)
	if env != nil {
		config, err := utils.ParseSqliteConfig(env.Value)

		if err != nil {
			return err
		}

		return utils.DumpSqliteConfig(sqlite3Path, config, out)
	}

	db := commands.FindOption("db", options)

	config := utils.SqliteConfig{
		DatabaseURL: "kyokusu.db",
	}

	if db != nil {
		config.DatabaseURL = db.Value
	}

	return utils.DumpSqliteConfig(sqlite3Path, config, out)
}


func getOutput(options []models.Option) (string, error) {
	output := commands.FindOption("o", options)
	format := commands.FindOption("f", options)

	path := constants.BaseOutput
	if output != nil && strings.TrimSpace(output.Value) != "" {
		path = strings.TrimSpace(output.Value)
	}

	formatName := constants.BaseFormat
	if format != nil && strings.TrimSpace(format.Value) != "" {
		formatName = strings.TrimPrefix(strings.TrimSpace(format.Value), ".")
	}

	if ext := filepath.Ext(path); ext != "" {
		extName := strings.TrimPrefix(ext, ".")
		if format == nil {
			formatName = extName
		} else if formatName != extName {
			return "", fmt.Errorf(
				"format %q conflicts with extension of %q",
				formatName, path,
			)
		}
	}

	if filepath.Ext(path) == "" {
		path += "." + formatName
	}

	return path, nil
}