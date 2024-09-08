package http

import (
	"context"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/ogen-go/ogen/middleware"
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/controller/http/handler"
	app "github.com/upikoth/starter-go/internal/generated/app"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/service"
)

type HTTP struct {
	logger    logger.Logger
	appServer *http.Server
}

func New(
	config *config.ControllerHTTP,
	loggerInstance logger.Logger,
	service *service.Service,
) (*HTTP, error) {
	handler := handler.New(loggerInstance, service)

	srv, err := app.NewServer(
		handler,
		app.WithErrorHandler(getAppErrorHandler(handler)),
		app.WithMiddleware(
			httpSentryMiddleware,
			logger.GetHTTPMiddleware(loggerInstance),
		),
	)

	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()

	mux.Handle("/api/", corsMiddleware(srv))

	app := http.FileServer(http.Dir("docs/app"))
	mux.Handle("/api/docs/app/", http.StripPrefix("/api/docs/app/", app))

	swaggerUI := http.FileServer(http.Dir("docs/swagger-ui"))
	mux.Handle("/api/docs/swagger-ui/", http.StripPrefix("/api/docs/swagger-ui/", swaggerUI))

	appServer := &http.Server{
		Addr:              ":" + config.Port,
		ReadHeaderTimeout: time.Minute,
		Handler:           mux,
	}

	return &HTTP{
		logger:    loggerInstance,
		appServer: appServer,
	}, nil
}

func (h *HTTP) Start() error {
	return h.appServer.ListenAndServe()
}

func (h *HTTP) Stop(ctx context.Context) error {
	return h.appServer.Shutdown(ctx)
}

func getAppErrorHandler(handler *handler.Handler) app.ErrorHandler {
	return func(context context.Context, w http.ResponseWriter, _ *http.Request, err error) {
		w.Header().Set("Content-Type", "application/json")

		res := handler.NewError(context, err)
		w.WriteHeader(res.StatusCode)

		bytes, _ := res.Response.MarshalJSON()
		_, _ = w.Write(bytes)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Baggage, Sentry-Trace, Authorization-Token")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func httpSentryMiddleware(req middleware.Request, next middleware.Next) (middleware.Response, error) {
	transaction := sentry.StartTransaction(
		req.Context,
		req.Raw.RequestURI,
		sentry.ContinueFromRequest(req.Raw),
	)

	ctx := transaction.Context()
	req.SetContext(ctx)

	res, errorResponse := next(req)

	if errorResponse != nil {
		sentry.CaptureEvent(&sentry.Event{
			Level:   sentry.LevelInfo,
			Message: errorResponse.Error(),
		})
	}

	transaction.Finish()

	return res, errorResponse
}
