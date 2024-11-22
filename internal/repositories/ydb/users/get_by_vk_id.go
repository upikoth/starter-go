package users

import (
	"context"

	"github.com/upikoth/starter-go/internal/models"
)

func (u *Users) GetByVkID(
	inputCtx context.Context,
	vkID string,
) (res *models.User, err error) {
	return u.getBy(inputCtx, fieldNameGetByVkID, vkID)
}
