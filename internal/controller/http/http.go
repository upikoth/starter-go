package http

import (
	"context"
	"net/http"
	"time"

	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/controller/http/handler"
	starter "github.com/upikoth/starter-go/internal/generated/starter"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/service"
)

type HTTP struct {
	logger        logger.Logger
	starterServer *http.Server
}

func New(
	config *config.ControllerHTTP,
	loggerInstance logger.Logger,
	service *service.Service,
) (*HTTP, error) {
	handler := handler.New(loggerInstance, service)

	srv, err := starter.NewServer(
		handler,
		starter.WithErrorHandler(getStarterErrorHandler(handler)),
		starter.WithMiddleware(logger.HTTPSentryMiddleware),
		starter.WithMiddleware(logger.GetHTTPMiddleware(loggerInstance)),
	)

	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()

	mux.Handle("/api/", corsMiddleware(srv))

	starter := http.FileServer(http.Dir("docs/starter"))
	mux.Handle("/api/docs/starter/", http.StripPrefix("/api/docs/starter/", starter))

	swaggerUI := http.FileServer(http.Dir("docs/swagger-ui"))
	mux.Handle("/api/docs/swagger-ui/", http.StripPrefix("/api/docs/swagger-ui/", swaggerUI))

	starterServer := &http.Server{
		Addr:              ":" + config.Port,
		ReadHeaderTimeout: time.Minute,
		Handler:           mux,
	}

	return &HTTP{
		logger:        loggerInstance,
		starterServer: starterServer,
	}, nil
}

func (h *HTTP) Start() error {
	return h.starterServer.ListenAndServe()
}

func (h *HTTP) Stop(ctx context.Context) error {
	return h.starterServer.Shutdown(ctx)
}

func getStarterErrorHandler(handler *handler.Handler) starter.ErrorHandler {
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
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Baggage, Sentry-Trace")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
