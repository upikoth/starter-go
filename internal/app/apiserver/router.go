package apiserver

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/upikoth/starter-go/internal/app/constants"
)

func (s *ApiServer) initRoutes() {
	s.router.SetTrustedProxies(nil)

	s.router.Use(formatResponse())

	s.router.GET("/api/v1/health", s.handler.V1.CheckHealth)

	s.router.POST("/api/v1/session", s.handler.V1.CreateSession)

	authorized := s.router.Use(checkAuthorization(s.config.JwtSecret))

	authorized.GET("/api/v1/session", s.handler.V1.GetSession)

	authorized.GET("/api/v1/users", s.handler.V1.GetUsers)

	authorized.POST("/api/v1/user", s.handler.V1.CreateUser)
	authorized.GET("/api/v1/users/:id", s.handler.V1.GetUser)
	authorized.DELETE("/api/v1/users/:id", s.handler.V1.DeleteUser)
	authorized.PATCH("/api/v1/users/:id", s.handler.V1.PatchUser)

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

func checkAuthorization(jwtSecret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken, err := c.Cookie("Authorization")

		if err != nil || jwtToken == "" {
			c.Set("responseCode", http.StatusForbidden)
			c.Set("responseErrorCode", constants.ErrUserNotAuthorized)
			c.Abort()
			return
		}

		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, err := token.Method.(*jwt.SigningMethodHMAC); !err {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.Set("responseCode", http.StatusForbidden)
			c.Set("responseErrorCode", constants.ErrUserNotAuthorized)
			c.Abort()
			return
		}

		c.Next()
	}
}
