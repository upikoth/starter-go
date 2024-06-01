package logger

import (
	"encoding/json"
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/ogen-go/ogen/middleware"
	"github.com/upikoth/starter-go/internal/config"
	loggerzerolog "github.com/upikoth/starter-go/internal/pkg/logger/logger-zerolog"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
}

func New() Logger {
	return loggerzerolog.New()
}

func InitSentry(
	config *config.ControllerHTTP,
	logger Logger,
) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.SentryDsn,
		Environment:      config.Environment,
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	})

	if err != nil {
		logger.Error(fmt.Sprintf("Не удалось инициализировать Sentry: %v\n", err))
	} else {
		logger.Debug("Sentry успешно запущена")
	}
}

func HTTPSentryMiddleware(req middleware.Request, next middleware.Next) (middleware.Response, error) {
	transaction := sentry.StartTransaction(req.Context, req.Raw.RequestURI, sentry.ContinueFromRequest(req.Raw))

	res, errorResponse := next(req)

	transaction.Finish()

	return res, errorResponse
}

func GetHTTPMiddleware(logger Logger) func(req middleware.Request, next middleware.Next) (middleware.Response, error) {
	return func(req middleware.Request, next middleware.Next) (middleware.Response, error) {
		reqBody, _ := json.Marshal(req.Body)
		logger.Debug(fmt.Sprintf("%s: %s", req.Raw.URL, string(reqBody)))

		res, errorResponse := next(req)

		resBody, _ := json.Marshal(res.Type)
		logger.Debug(fmt.Sprintf("%s: %s", req.Raw.URL, string(resBody)))

		return res, errorResponse
	}
}
