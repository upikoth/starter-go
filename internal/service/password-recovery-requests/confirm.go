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
) (models.Session, error) {
	span := sentry.StartSpan(inputCtx, "Service: PasswordRecoveryRequests.Confirm")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	passwordRecoveryRequest, err := p.
		repository.
		YdbStarter.
		PasswordRecoveryRequests.
		GetByToken(ctx, params.ConfirmationToken)

	if err != nil {
		sentry.CaptureException(err)
		return models.Session{}, &models.Error{
			Code:        models.ErrorCodePasswordRecoveryRequestYdbStarterCheckConfirmationToken,
			Description: err.Error(),
		}
	}

	if passwordRecoveryRequest.ID == "" {
		return models.Session{}, &models.Error{
			Code:        models.ErrorCodePasswordRecoveryRequestPasswordRecoveryRequestNotFound,
			Description: "Заявка на восстановление пароля с переданным токеном не найдена",
			StatusCode:  http.StatusBadRequest,
		}
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(params.NewPassword), bcrypt.DefaultCost)

	if err != nil {
		sentry.CaptureException(err)
		return models.Session{}, &models.Error{
			Code:        models.ErrorCodePasswordRecoveryRequestGeneratePasswordHash,
			Description: err.Error(),
		}
	}

	user, err := p.repository.YdbStarter.Users.GetByEmail(ctx, passwordRecoveryRequest.Email)

	if err != nil {
		sentry.CaptureException(err)
		return models.Session{}, &models.Error{
			Code:        models.ErrorCodePasswordRecoveryRequestFindUserByEmail,
			Description: err.Error(),
		}
	}

	user.PasswordHash = string(passwordHash)

	updatedUser, err := p.
		repository.
		YdbStarter.
		PasswordRecoveryRequestsAndUsers.
		DeletePasswordRecoveryRequestAndUpdateUser(
			ctx,
			passwordRecoveryRequest,
			user,
		)

	if err != nil {
		sentry.CaptureException(err)
		return models.Session{}, &models.Error{
			Code:        models.ErrorCodePasswordRecoveryRequestUpdateUserPassword,
			Description: err.Error(),
		}
	}

	session := models.Session{
		ID:     uuid.New().String(),
		UserID: updatedUser.ID,
		Token:  uuid.New().String(),
	}

	createdSession, err := p.repository.YdbStarter.Sessions.Create(ctx, session)

	if err != nil {
		sentry.CaptureException(err)
		return models.Session{}, &models.Error{
			Code:        models.ErrorCodePasswordRecoveryRequestCreateSession,
			Description: err.Error(),
		}
	}

	return createdSession, nil
}
