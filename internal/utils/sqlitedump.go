package utils

import "os/exec"

func DumpSqliteConfig(sqlite3Path string, config SqliteConfig, outputPath string) error {
	cmd := exec.Command(
		sqlite3Path,
		config.DatabaseURL,
		".output "+outputPath,
		".dump",
	)
	return cmd.Run()
}
