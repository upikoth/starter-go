package http

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/upikoth/starter-go/internal/controller/http/v1"
	"github.com/upikoth/starter-go/internal/pkg/logger"
)

type HTTP struct {
	v1     *v1.HandlerV1
	router *gin.Engine
	config *Config
	logger logger.Logger
}

func New(c *Config, logger logger.Logger) *HTTP {
	gin.SetMode(gin.ReleaseMode)

	return &HTTP{
		v1:     v1.New(logger),
		router: gin.New(),
		config: c,
		logger: logger,
	}
}

func (h *HTTP) Start() error {
	h.startRouting()

	return h.router.Run(":" + h.config.Port)
}
