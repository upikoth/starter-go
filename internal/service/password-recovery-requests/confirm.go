package passwordrecoveryrequests

import (
	"context"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	"github.com/upikoth/starter-go/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (p *PasswordRecoveryRequests) Confirm(
	inputCtx context.Context,
	params models.PasswordRecoveryRequestConfirmParams,
) (*models.Session, error) {
	span := sentry.StartSpan(inputCtx, "Service: PasswordRecoveryRequests.Confirm")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	passwordRecoveryRequest, err := p.
		repository.
		YDB.
		PasswordRecoveryRequests.
		GetByToken(ctx, params.ConfirmationToken)

	if err != nil {
		sentry.CaptureException(err)
		return nil, &models.Error{
			Code:        models.ErrorCodePasswordRecoveryRequestYdbCheckConfirmationToken,
			Description: err.Error(),
		}
	}

	if passwordRecoveryRequest.ID == "" {
		return nil, &models.Error{
			Code:        models.ErrorCodePasswordRecoveryRequestPasswordRecoveryRequestNotFound,
			Description: "Password recovery request with transferred token not found",
			StatusCode:  http.StatusBadRequest,
		}
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(params.NewPassword), bcrypt.DefaultCost)

	if err != nil {
		sentry.CaptureException(err)
		return nil, &models.Error{
			Code:        models.ErrorCodePasswordRecoveryRequestGeneratePasswordHash,
			Description: err.Error(),
		}
	}

	user, err := p.repository.YDB.Users.GetByEmail(ctx, passwordRecoveryRequest.Email)

	if err != nil {
		sentry.CaptureException(err)
		return nil, &models.Error{
			Code:        models.ErrorCodePasswordRecoveryRequestFindUserByEmail,
			Description: err.Error(),
		}
	}

	user.PasswordHash = string(passwordHash)

	updatedUser, err := p.
		repository.
		YDB.
		PasswordRecoveryRequestsAndUsers.
		Delete(
			ctx,
			*passwordRecoveryRequest,
			*user,
		)

	if err != nil {
		sentry.CaptureException(err)
		return nil, &models.Error{
			Code:        models.ErrorCodePasswordRecoveryRequestUpdateUserPassword,
			Description: err.Error(),
		}
	}

	session := models.Session{
		ID:       uuid.New().String(),
		UserID:   updatedUser.ID,
		UserRole: updatedUser.Role,
		Token:    uuid.New().String(),
	}

	createdSession, err := p.repository.YDB.Sessions.Create(ctx, session)

	if err != nil {
		sentry.CaptureException(err)
		return nil, &models.Error{
			Code:        models.ErrorCodePasswordRecoveryRequestCreateSession,
			Description: err.Error(),
		}
	}

	return createdSession, nil
}
