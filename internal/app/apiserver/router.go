package apiserver

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/upikoth/starter-go/internal/app/constants"
)

func (s *ApiServer) initRoutes() {
	s.router.SetTrustedProxies(nil)

	s.router.Use(formatResponse())

	s.router.GET("/api/v1/health", s.handler.V1.CheckHealth)
	s.router.GET("/api/v1/users", s.handler.V1.GetUsers)

	s.router.NoRoute(func(c *gin.Context) {
		c.Set("responseCode", http.StatusNotFound)
		c.Set("responseErrorCode", constants.ErrRouteNotFound)
	})
}

func formatResponse() gin.HandlerFunc {
	type ResponseError struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	}

	type Response struct {
		Success bool           `json:"success"`
		Data    interface{}    `json:"data"`
		Error   *ResponseError `json:"error,omitempty"`
	}

	return func(c *gin.Context) {
		c.Next()

		response := Response{}

		code, isCodeExist := c.Get("responseCode")
		data, isDataExist := c.Get("responseData")
		errorCode, isErorrCodeExist := c.Get("responseErrorCode")

		if !isCodeExist {
			code = http.StatusOK
		}

		if !isDataExist {
			data = map[string]string{}
		}

		if isErorrCodeExist {
			response.Success = false
			response.Error = &ResponseError{
				Code:        fmt.Sprintf("%v", errorCode),
				Description: constants.ErrDescriptionByCode[errorCode.(error)],
			}
		} else {
			response.Success = true
		}

		response.Data = data

		c.JSON(code.(int), response)
	}
}
