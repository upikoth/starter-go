package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/upikoth/starter-go/docs"
	"github.com/upikoth/starter-go/internal/constants"
)

var (
	ResponseCode         = "responseCode"
	ResponseData         = "responseData"
	ResponseErrorCode    = "responseErrorCode"
	ResponseErrorDetails = "responseErrorDetails"
)

func (h *HTTP) startRouting() {
	docs.SwaggerInfo.Schemes = []string{}
	h.router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	proxiesErr := h.router.SetTrustedProxies(nil)

	if proxiesErr != nil {
		log.Println(proxiesErr)
	}

	h.router.Use(gin.Recovery())
	h.router.Use(loggingMiddleware(h.logger))
	h.router.Use(corsMiddleware())
	h.router.Use(formatResponse())

	h.router.GET("/api/v1/health", h.v1.CheckHealth)

	h.router.POST("/api/v1/registrations", h.v1.CreateRegistration)
	h.router.PATCH("/api/v1/registrations", h.v1.ConfirmRegistration)

	h.router.GET("/api/v1/users", h.v1.GetUsers)
	h.router.GET("/api/v1/users/:id", h.v1.GetUser)

	h.router.GET("/api/v1/sessions", h.v1.GetSessions)
	h.router.POST("/api/v1/sessions", h.v1.CreateSession)

	h.router.NoRoute(func(c *gin.Context) {
		c.Set(ResponseCode, http.StatusNotFound)
		c.Set(ResponseErrorCode, constants.ErrRouteNotFound)
	})
}
