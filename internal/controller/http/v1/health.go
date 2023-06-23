package v1

import (
	"github.com/gin-gonic/gin"
)

// CheckHealth godoc
// @Summary      Проверка работоспособности сервера
// @Tags         health
// @Success      200  {object}  http.ResponseSuccess{data=interface{}}
// @Router       /api/v1/health [get].
func (h *HandlerV1) CheckHealth(_ *gin.Context) {}
