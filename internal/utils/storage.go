package utils

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"

	"github.com/joho/godotenv"

)

type PostgresConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	DatabaseName string

	DatabaseURL string
	SSLMode     string
}

type SqliteConfig struct {
	DatabaseURL string
}

func ParsePostgresConfig(env string) (PostgresConfig, error) {
	if err := loadEnv(env); err != nil {
		return PostgresConfig{}, err
	}
	return parsePostgresConfig()
}

func ParseSqliteConfig(env string) (SqliteConfig, error) {
	if err := loadEnv(env); err != nil {
		return SqliteConfig{}, err
	}
	return parseSqliteConfig()
}

func parsePostgresConfig() (PostgresConfig, error) {
	cfg := PostgresConfig{
		Host:         getEnv("DATABASE_HOST", "postgres"),
		Port:         getEnv("DATABASE_PORT", "5432"),
		User:         getEnv("DATABASE_USER", "postgres"),
		Password:     getEnv("DATABASE_PASSWORD", ""),
		DatabaseName: getEnv("DATABASE_NAME", ""),
		DatabaseURL:  getEnv("DATABASE_URL", ""),
		SSLMode:      getEnv("DATABASE_SSL_MODE", "disable"),
	}
	
	if cfg.DatabaseURL != "" {
		return cfg, nil
	}

	if cfg.DatabaseName == "" {
		return PostgresConfig{}, errors.New("DB_NAME is required")
	}

	if cfg.Password == "" {
		return PostgresConfig{}, errors.New("DATABASE_PASSWORD is required")
	}

	u := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.User, cfg.Password),
		Host:   net.JoinHostPort(cfg.Host, cfg.Port),
		Path:   cfg.DatabaseName,
	}
	
	cfg.DatabaseURL = u.String() + "?sslmode=" + cfg.SSLMode

	return cfg, nil
}

func parseSqliteConfig() (SqliteConfig, error) {
	cfg := SqliteConfig{DatabaseURL: getEnv("DATABASE_URL", "")}

	if cfg.DatabaseURL == "" {
		return SqliteConfig{}, errors.New("DATABASE_URL is required")
	}

	return cfg, nil
}

func loadEnv(path string) error {
	if path == "" {
		return nil
	}

	err := godotenv.Load(path)
	if err == nil || errors.Is(err, os.ErrNotExist) {
		return nil
	}

	return fmt.Errorf("load env file %s: %w", path, err)
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}