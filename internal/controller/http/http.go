package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/ogen-go/ogen/middleware"
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
	logger logger.Logger,
	service *service.Service,
) (*HTTP, error) {
	handler := handler.New(logger, service)

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.SentryDsn,
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	})

	if err != nil {
		logger.Error(fmt.Sprintf("Sentry initialization failed: %v\n", err))
	}

	srv, err := starter.NewServer(
		handler,
		starter.WithErrorHandler(getStarterErrorHandler(handler)),
		starter.WithMiddleware(func(req middleware.Request, next middleware.Next) (middleware.Response, error) {
			transaction := sentry.StartTransaction(req.Context, req.Raw.RequestURI, sentry.ContinueFromRequest(req.Raw))
			res, errorResponse := next(req)
			transaction.Finish()
			return res, errorResponse
		}),
	)

	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()

	mux.Handle("/api/", srv)

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
		logger:        logger,
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
