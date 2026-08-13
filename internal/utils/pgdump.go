package utils

import (
	"os/exec"
)

func DumpPostgresConfig(pgDumpPath string, config PostgresConfig, outputPath string) error {
	var cmd *exec.Cmd

	if config.DatabaseURL == "" {
		cmd = exec.Command(pgDumpPath, "-h", config.Host, "-p", config.Port, "-U", config.User, "-d", config.DatabaseName, "-f", outputPath)
	} else {
		cmd = exec.Command(pgDumpPath, "-d", config.DatabaseURL, "-f", outputPath)
	}

	err := cmd.Run()
	
	if err != nil {
		return err
	}
	
	return nil
}
