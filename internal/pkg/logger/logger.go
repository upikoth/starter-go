package logger

import (
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

func New() Logger {
	config, configErr := newConfig()

	var logger *zap.Logger

	sentryOption, sentryError := getZapSentryOptions(config)

	if configErr != nil || config.AppEnv != "development" {
		logger = zap.Must(zap.NewProduction(sentryOption))
	} else {
		logger = zap.Must(zap.NewDevelopment(sentryOption))
	}

	if configErr != nil {
		logger.Error(configErr.Error())
	}

	if sentryError != nil {
		logger.Fatal(sentryError.Error())
	}

	sugar := logger.Sugar()

	return sugar
}

func getZapSentryOptions(config *config) (zap.Option, error) {
	sentryErr := sentry.Init(sentry.ClientOptions{
		TracesSampleRate: 1.0,
		Dsn:              config.SentryDsn,
	})

	if sentryErr != nil {
		return nil, sentryErr
	}

	option := zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.RegisterHooks(core, func(entry zapcore.Entry) error {
			if entry.Level < zap.WarnLevel {
				return nil
			}

			defer sentry.Flush(time.Second)

			event := sentry.NewEvent()
			event.Message = entry.Message
			event.Level = getSentryLogLevelFromZapLogLevel(entry.Level)
			event.Timestamp = entry.Time
			event.Environment = config.AppEnv

			sentry.CaptureEvent(event)

			return nil
		})
	})

	return option, nil
}

func getSentryLogLevelFromZapLogLevel(level zapcore.Level) sentry.Level {
	switch level {
	case zapcore.DebugLevel:
		return sentry.LevelDebug
	case zapcore.InfoLevel:
		return sentry.LevelInfo
	case zapcore.WarnLevel:
		return sentry.LevelWarning
	case zapcore.ErrorLevel:
		return sentry.LevelError
	case zapcore.FatalLevel:
		return sentry.LevelFatal

	case zapcore.DPanicLevel:
		return sentry.LevelInfo
	case zapcore.PanicLevel:
		return sentry.LevelInfo
	case zapcore.InvalidLevel:
		return sentry.LevelInfo
	default:
		return sentry.LevelInfo
	}
}

type config struct {
	AppEnv    string `envconfig:"APP_ENV" required:"true"`
	SentryDsn string `envconfig:"SENTRY_DSN" required:"true"`
}

func newConfig() (*config, error) {
	config := &config{}

	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}

	return config, nil
}
