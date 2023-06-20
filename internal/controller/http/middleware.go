package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/upikoth/starter-go/internal/constants"
)

type ResponseErrorField struct {
	Code            string `json:"code"`
	CodeDescription string `json:"codeDescription"`
	Details         string `json:"details"`
}

type ResponseSuccess struct {
	Success bool        `json:"success" default:"true"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Success bool                `json:"success" default:"false"`
	Error   *ResponseErrorField `json:"error"`
}

func formatResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		code, isCodeExist := c.Get(ResponseCode)
		data, isDataExist := c.Get(ResponseData)
		errorCode, isErorrCodeExist := c.Get(ResponseErrorCode)
		errorDetails, isErorrDetailsExist := c.Get(ResponseErrorDetails)

		if !isCodeExist {
			code = http.StatusOK
		}

		if !isDataExist {
			data = map[string]string{}
		}

		if !isErorrDetailsExist {
			errorDetails = ""
		}

		isResponseSuccess := !isErorrCodeExist

		if isResponseSuccess {
			response := ResponseSuccess{
				Success: true,
				Data:    data,
			}

			c.JSON(code.(int), response)
		} else {
			response := ResponseError{
				Success: false,
				Error: &ResponseErrorField{
					Code:            fmt.Sprintf("%v", errorCode),
					CodeDescription: constants.ErrDescriptionByCode[errorCode.(error)],
					Details:         fmt.Sprintf("%v", errorDetails),
				},
			}
			c.JSON(code.(int), response)
		}
	}
}
