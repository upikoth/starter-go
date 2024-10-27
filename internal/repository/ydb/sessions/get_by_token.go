package sessions

import (
	"context"

	"github.com/upikoth/starter-go/internal/models"
)

func (s *Sessions) GetByToken(
	inputCtx context.Context,
	token string,
) (res *models.SessionWithUserRole, err error) {
	return s.getBy(inputCtx, fieldNameGetByToken, token)
}
