package http

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/upikoth/starter-go/internal/controller/http/v1"
)

type HTTP struct {
	v1     *v1.HandlerV1
	router *gin.Engine
	config *Config
}

func New(c *Config) *HTTP {
	return &HTTP{
		v1:     v1.New(),
		router: gin.Default(),
		config: c,
	}
}

func (h *HTTP) Start() error {
	h.startRouting()

	return h.router.Run(":" + h.config.Port)
}
