package apiserver

import (
	"github.com/gin-gonic/gin"
	"github.com/upikoth/starter-go/internal/app/handler"
)

type ApiServer struct {
	config  *Config
	router  *gin.Engine
	handler *handler.Handler
}

func New(config *Config) *ApiServer {
	handler := handler.New()

	return &ApiServer{
		config:  config,
		router:  gin.Default(),
		handler: handler,
	}
}

func (s *ApiServer) Start() error {
	s.initRoutes()

	return s.router.Run(":" + s.config.Port)
}
