package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/upikoth/starter-go/internal/app"
)

// @title   Starter API.
func main() {
	// Чтение .env файла нужно только при локальной разработке.
	// В других случаях значения переменных окружения уже должны быть установлены.
	_ = godotenv.Load()

	config, configErr := app.NewConfig()
	if configErr != nil {
		log.Fatal(configErr)
	}

	server := app.New(config)

	if serverErr := server.Start(); serverErr != nil {
		log.Fatal(serverErr)
	}
}
