package users

import (
	"context"

	"github.com/upikoth/starter-go/internal/models"
)

func (u *Users) GetByMailRuID(
	inputCtx context.Context,
	mailRuID string,
) (res *models.User, err error) {
	return u.getBy(inputCtx, fieldNameGetByMailRuID, mailRuID)
}
