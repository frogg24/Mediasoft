package config

import "os"

type Config struct {
	PG   string
	Port string
}

func NewConfig() *Config {
	return &Config{
		PG:   os.Getenv("DATABASE_URL"),
		Port: os.Getenv("PORT"),
	}
}
