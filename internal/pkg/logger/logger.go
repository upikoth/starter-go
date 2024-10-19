package logger

import (
	"encoding/json"
	"fmt"

	"github.com/ogen-go/ogen/middleware"
	loggerzerolog "github.com/upikoth/starter-go/internal/pkg/logger/logger-zerolog"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
	SetPrettyOutputToConsole()
}

func New() Logger {
	return loggerzerolog.New()
}

func GetHTTPMiddleware(logger Logger) func(req middleware.Request, next middleware.Next) (middleware.Response, error) {
	return func(req middleware.Request, next middleware.Next) (middleware.Response, error) {
		reqBody, _ := json.Marshal(req.Body)
		logger.Debug(fmt.Sprintf("Request: %s: %s", req.Raw.URL, string(reqBody)))

		res, errorResponse := next(req)

		if errorResponse != nil {
			logger.Info(errorResponse.Error())
		}

		resBody, _ := json.Marshal(res.Type)
		logger.Debug(fmt.Sprintf("Response: %s: %s", req.Raw.URL, string(resBody)))

		return res, errorResponse
	}
}
