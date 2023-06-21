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

	h.router.Use(corsMiddleware())
	h.router.Use(formatResponse())

	h.router.GET("/api/v1/health", h.v1.CheckHealth)

	h.router.NoRoute(func(c *gin.Context) {
		c.Set(ResponseCode, http.StatusNotFound)
		c.Set(ResponseErrorCode, constants.ErrRouteNotFound)
	})
}
