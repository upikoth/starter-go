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

	starterHandler, err := starter.NewServer(handler)

	if err != nil {
		return nil, err
	}

	starterServer := &http.Server{
		Addr:              ":" + config.Port,
		ReadHeaderTimeout: time.Minute,
		Handler:           starterHandler,
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
