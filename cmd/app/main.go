package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/upikoth/starter-go/internal/app"
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
)

func main() {
	// Чтение .env файла нужно только при локальной разработке.
	// В других случаях значения переменных окружения уже должны быть установлены.
	// Поэтому ошибку загрузки файла обрабатывать не нужно.
	_ = godotenv.Load()
	loggerInstance := logger.New()

	config, err := config.New()
	if err != nil {
		loggerInstance.Fatal(err.Error())
	}

	if config.Environment == "development" {
		loggerInstance.SetPrettyOutputToConsole()
	}

	logger.InitSentry(
		&config.Controller.HTTP,
		loggerInstance,
	)

	app, err := app.New(config, loggerInstance)
	if err != nil {
		loggerInstance.Fatal(err.Error())
	}

	go func() {
		loggerInstance.Info("Запуск приложения")

		if appErr := app.Start(); !errors.Is(appErr, http.ErrServerClosed) {
			loggerInstance.Fatal(appErr.Error())
		}

		loggerInstance.Info("Приложение перестало принимать новые запросы")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	timeToStopAppInSeconds := 10
	shutdownCtx, shutdownRelease := context.WithTimeout(
		context.Background(),
		time.Duration(timeToStopAppInSeconds)*time.Second,
	)
	defer shutdownRelease()

	if stopErr := app.Stop(shutdownCtx); stopErr != nil {
		loggerInstance.Fatal(fmt.Sprintf("Не удалось корректно остановить сервер, ошибка: %v", stopErr))
	}

	loggerInstance.Info("Приложение остановлено")
}
