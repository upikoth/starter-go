package apiserver

import (
	"os"
)

type Config struct {
	Port string
}

func NewConfig() *Config {
	return &Config{
		Port: os.Getenv("APP_PORT"),
	}
}
