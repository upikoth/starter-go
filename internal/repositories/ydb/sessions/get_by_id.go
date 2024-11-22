package sessions

import (
	"context"

	"github.com/upikoth/starter-go/internal/models"
)

func (s *Sessions) GetByID(
	inputCtx context.Context,
	id string,
) (res *models.SessionWithUserRole, err error) {
	return s.getBy(inputCtx, fieldNameGetByID, id)
}
