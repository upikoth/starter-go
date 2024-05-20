package http

import (
	"context"
	"net/http"
	"time"

	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/controller/http/handler"
	starter "github.com/upikoth/starter-go/internal/generated/starter"
	"github.com/upikoth/starter-go/internal/pkg/logger"
)

type HTTP struct {
	logger        logger.Logger
	starterServer *http.Server
}

func New(config *config.ControllerHTTP, logger logger.Logger) (*HTTP, error) {
	handler := handler.New(logger)

	srv, err := starter.NewServer(handler)

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
