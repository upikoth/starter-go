package store

import "os"

type Config struct {
	DatabaseName     string
	DatabaseAddr     string
	DatabaseUser     string
	DatabasePassword string
}

func NewConfig() *Config {
	return &Config{
		DatabaseName:     os.Getenv("DATABASE_NAME"),
		DatabaseAddr:     os.Getenv("DATABASE_ADDR"),
		DatabaseUser:     os.Getenv("DATABASE_USER"),
		DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
	}
}
