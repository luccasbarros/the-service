package main

import "os"

type DatabaseConfig struct {
	URI          string
	MigrationDir string
}

type Config struct {
	Database DatabaseConfig
}

func NewConfig() *Config {
	return &Config{
		Database: DatabaseConfig{
			URI:          os.Getenv("DNS"),
			MigrationDir: os.Getenv("MIGRATIONS_DIR"),
		},
	}
}
