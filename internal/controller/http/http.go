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
	logger logger.Logger,
	service *service.Service,
) (*HTTP, error) {
	handler := handler.New(logger, service)

	srv, err := starter.NewServer(
		handler,
		starter.WithErrorHandler(getStarterErrorHandler(handler)),
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
