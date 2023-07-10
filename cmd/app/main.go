package main

import (
	"github.com/joho/godotenv"
	"github.com/upikoth/starter-go/internal/app"
	"github.com/upikoth/starter-go/internal/pkg/logger"
)

// @title   Starter API.
func main() {
	// Чтение .env файла нужно только при локальной разработке.
	// В других случаях значения переменных окружения уже должны быть установлены.
	// Поэтому ошибку загрузки файла обрабатывать не нужно.
	_ = godotenv.Load()
	logger := logger.New()

	config, configErr := app.NewConfig()
	if configErr != nil {
		logger.Fatal(configErr)
	}

	app := app.New(config, logger)

	logger.Info("Запуск приложения")
	if appErr := app.Start(); appErr != nil {
		logger.Fatal(appErr)
	}
}
