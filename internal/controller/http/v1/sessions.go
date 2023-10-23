package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/upikoth/starter-go/internal/constants"
	"github.com/upikoth/starter-go/internal/model"
)

type sessionsResponseData struct {
	Sessions []model.Session `json:"sessions"`
}

// GetSessions godoc
// @Summary      Получение списка сессий
// @Tags         sessions
// @Produce      json
// @Success      200  {object}  http.ResponseSuccess{data=sessionsResponseData}
// @Failure      500  {object}  http.ResponseError
// @Router       /api/v1/sessions [get].
func (h *HandlerV1) GetSessions(c *gin.Context) {
	sessions, err := h.service.Sessions.GetAll()

	if err != nil {
		c.Set("ResponseCode", http.StatusInternalServerError)
		c.Set("responseErrorCode", constants.ErrDbError)
		c.Set("responseErrorDetails", err)
		return
	}

	c.Set("responseData", sessionsResponseData{
		Sessions: sessions,
	})
}
