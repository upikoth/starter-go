package users

import (
	"context"

	"github.com/upikoth/starter-go/internal/models"
)

func (u *Users) GetByEmail(
	inputCtx context.Context,
	email string,
) (res *models.User, err error) {
	return u.getBy(inputCtx, fieldNameGetByEmail, email)
}
