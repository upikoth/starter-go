package ydbmodels

import "github.com/upikoth/starter-go/internal/models"

type PasswordRecoveryRequest struct {
	ID                string `gorm:"primarykey"`
	Email             string
	ConfirmationToken string
}

func NewYDBPasswordRecoveryRequestModel(registration models.PasswordRecoveryRequest) PasswordRecoveryRequest {
	return PasswordRecoveryRequest{
		ID:                registration.ID,
		Email:             registration.Email,
		ConfirmationToken: registration.ConfirmationToken,
	}
}

func (r *PasswordRecoveryRequest) FromYDBModel() models.PasswordRecoveryRequest {
	return models.PasswordRecoveryRequest{
		ID:                r.ID,
		Email:             r.Email,
		ConfirmationToken: r.ConfirmationToken,
	}
}
