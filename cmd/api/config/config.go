package config

import "os"

type DBConfig struct {
	User     string
	Password string
	Name     string
	Host     string
	Port     string
	Secret   string
}

var Env = DBConfig{
	User:     os.Getenv("DB_USER"),
	Password: os.Getenv("DB_PASSWORD"),
	Name:     os.Getenv("DB_NAME"),
	Host:     os.Getenv("DB_HOST"),
	Port:     os.Getenv("DB_PORT"),
	Secret:   os.Getenv("JWT_SECRET"),
}
