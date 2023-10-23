package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/upikoth/starter-go/internal/constants"
	"github.com/upikoth/starter-go/internal/pkg/logger"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

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

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

type httpRequestLoggingInfo struct {
	URL         string      `json:"url"`
	RequestBody interface{} `json:"requestBody"`
}

type httpResponseLoggingInfo struct {
	URL          string      `json:"url"`
	ResponseCode int         `json:"responseCode"`
	ResponseBody interface{} `json:"responseBody"`
}

func loggingMiddleware(logger logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &bodyLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = blw

		requestBodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBodyBytes))
		var requestBodyJSON interface{}
		_ = json.Unmarshal(requestBodyBytes, &requestBodyJSON)
		requestLogInfo := httpRequestLoggingInfo{
			URL:         c.Request.RequestURI,
			RequestBody: requestBodyJSON,
		}

		requestLogInfoString, _ := json.Marshal(requestLogInfo)

		logger.Info(fmt.Sprintf("Запрос %s", requestLogInfoString))

		c.Next()

		responseBody := map[string]interface{}{}
		_ = json.Unmarshal(blw.body.Bytes(), &responseBody)

		responseLogInfo := httpResponseLoggingInfo{
			URL:          c.Request.RequestURI,
			ResponseCode: c.Writer.Status(),
			ResponseBody: responseBody,
		}

		responseLogInfoString, _ := json.Marshal(responseLogInfo)

		logger.Info(fmt.Sprintf("Ответ %s", responseLogInfoString))
	}
}
