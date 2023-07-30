package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/upikoth/starter-go/internal/constants"
	"github.com/upikoth/starter-go/internal/model"
)

type getUsersResponseData struct {
	Users []model.User `json:"users"`
}

// GetUsers godoc
// @Summary      Получение списка пользователей
// @Tags         users
// @Success      200  {object}  http.ResponseSuccess{data=getUsersResponseData}
// @Router       /api/v1/users [get].
func (h *HandlerV1) GetUsers(c *gin.Context) {
	users, err := h.service.GetUsers()

	if err != nil {
		c.Set("ResponseCode", http.StatusInternalServerError)
		c.Set("responseErrorCode", constants.ErrDbError)
		c.Set("responseErrorDetails", err)
		return
	}

	c.Set("responseData", getUsersResponseData{
		Users: users,
	})
}
