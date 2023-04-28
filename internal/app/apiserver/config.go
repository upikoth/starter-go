package apiserver

import (
	"os"
)

type Config struct {
	Port      string
	JwtSecret []byte
}

func NewConfig() *Config {
	return &Config{
		Port:      os.Getenv("APP_PORT"),
		JwtSecret: []byte(os.Getenv("JWT_SECRET")),
	}
}
