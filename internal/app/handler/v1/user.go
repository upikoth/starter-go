package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/upikoth/starter-go/internal/app/model"
)

type getUsersResponseDataV1 struct {
	Users []model.User `json:"users"`
}

func (h *HandlerV1) GetUsers(c *gin.Context) {
	users, err := h.store.GetUsers()

	if err != nil {
		c.Set("responseErrorCode", err)
		return
	}

	responseData := getUsersResponseDataV1{users}
	c.Set("responseData", responseData)
}
