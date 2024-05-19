package http

import (
	"github.com/gin-gonic/gin"
	"github.com/upikoth/starter-go/internal/config"
	v1 "github.com/upikoth/starter-go/internal/controller/http/v1"
	"github.com/upikoth/starter-go/internal/pkg/logger"
)

type HTTP struct {
	v1     *v1.HandlerV1
	router *gin.Engine
	config *config.ControllerHTTP
	logger logger.Logger
}

func New(config *config.ControllerHTTP, logger logger.Logger) *HTTP {
	gin.SetMode(gin.ReleaseMode)

	return &HTTP{
		v1:     v1.New(logger),
		router: gin.New(),
		config: config,
		logger: logger,
	}
}

func (h *HTTP) Start() error {
	h.startRouting()

	return h.router.Run(":" + h.config.Port)
}
