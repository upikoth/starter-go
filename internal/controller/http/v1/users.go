package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/upikoth/starter-go/internal/constants"
	"github.com/upikoth/starter-go/internal/model"
)

type usersResponseData struct {
	Users []model.User `json:"users"`
}

// GetUsers godoc
// @Summary      Получение списка пользователей
// @Tags         users
// @Produce      json
// @Success      200  {object}  http.ResponseSuccess{data=usersResponseData}
// @Failure      500  {object}  http.ResponseError
// @Router       /api/v1/users [get].
func (h *HandlerV1) GetUsers(c *gin.Context) {
	users, err := h.service.GetUsers()

	if err != nil {
		c.Set("ResponseCode", http.StatusInternalServerError)
		c.Set("responseErrorCode", constants.ErrDbError)
		c.Set("responseErrorDetails", err)
		return
	}

	c.Set("responseData", usersResponseData{
		Users: users,
	})
}

type userRequestData struct {
	ID int `json:"id" uri:"id" binding:"required"`
}

type userResponseData struct {
	User model.User `json:"user"`
}

// GetUser godoc
// @Summary      Получение пользователя по id
// @Tags         users
// @Produce      json
// @Param        id  path  int  true  "Id пользователя"
// @Success      200  {object}  http.ResponseSuccess{data=userResponseData}
// @Failure      400  {object}  http.ResponseError
// @Failure      500  {object}  http.ResponseError
// @Router       /api/v1/users/{id} [get].
func (h *HandlerV1) GetUser(c *gin.Context) {
	data := userRequestData{}
	err := c.BindUri(&data)

	if err != nil {
		c.Set("ResponseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrRequestValidation)
		c.Set("responseErrorDetails", err)
		return
	}

	user, err := h.service.GetUser(data.ID)

	if err != nil {
		c.Set("ResponseCode", http.StatusInternalServerError)
		c.Set("responseErrorCode", constants.ErrDbError)
		c.Set("responseErrorDetails", err)
		return
	}

	c.Set("responseData", userResponseData{
		User: user,
	})
}
