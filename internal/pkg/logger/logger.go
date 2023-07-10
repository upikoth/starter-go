package logger

import (
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

func New() Logger {
	config, configErr := newConfig()

	var logger *zap.Logger

	if configErr != nil || config.AppEnv != "development" {
		logger = zap.Must(zap.NewProduction())
	} else {
		logger = zap.Must(zap.NewDevelopment())
	}

	if configErr != nil {
		logger.Error(configErr.Error())
	}

	sugar := logger.Sugar()

	return sugar
}

type Config struct {
	AppEnv string `envconfig:"APP_ENV" required:"true"`
}

func newConfig() (*Config, error) {
	config := &Config{}

	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}

	return config, nil
}
