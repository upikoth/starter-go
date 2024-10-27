package passwordrecoveryrequests

import (
	"context"

	"github.com/upikoth/starter-go/internal/models"
)

func (p *PasswordRecoveryRequests) GetByToken(
	inputCtx context.Context,
	confirmationToken string,
) (res *models.PasswordRecoveryRequest, err error) {
	return p.getBy(inputCtx, fieldNameGetByConfrimationToken, confirmationToken)
}
