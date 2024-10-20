package registrations

import (
	"context"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	"github.com/upikoth/starter-go/internal/models"
	"github.com/upikoth/starter-go/internal/repository/ydb"

	"golang.org/x/crypto/bcrypt"
)

func (r *Registrations) Confirm(
	inputCtx context.Context,
	params models.RegistrationConfirmParams,
) (*models.Session, error) {
	span := sentry.StartSpan(inputCtx, "Service: Registrations.Confirm")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	registration, err := r.repository.YDB.Registrations.GetByToken(ctx, params.ConfirmationToken)

	if err != nil {
		sentry.CaptureException(err)
		return nil, &models.Error{
			Code:        models.ErrorCodeRegistrationYdbCheckConfirmationToken,
			Description: err.Error(),
		}
	}

	if registration.ID == "" {
		return nil, &models.Error{
			Code:        models.ErrorCodeRegistrationRegistrationNotFound,
			Description: "Registration with transferred token not found",
			StatusCode:  http.StatusBadRequest,
		}
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)

	if err != nil {
		sentry.CaptureException(err)
		return nil, &models.Error{
			Code:        models.ErrorCodeRegistrationGeneratePasswordHash,
			Description: err.Error(),
		}
	}

	newUser := models.User{
		ID:           uuid.New().String(),
		Email:        registration.Email,
		PasswordHash: string(passwordHash),
		Role:         models.UserRoleUser,
	}

	var createdUser *models.User
	err = r.repository.YDB.Transaction(func(ydbTx *ydb.YDB) error {
		dbErr := ydbTx.Registrations.Delete(ctx, registration.ID)

		if dbErr != nil {
			return dbErr
		}

		dbRes, dbErr := ydbTx.Users.Create(ctx, &newUser)
		createdUser = dbRes

		if dbErr != nil {
			return dbErr
		}

		return nil
	})

	if err != nil {
		sentry.CaptureException(err)
		return nil, &models.Error{
			Code:        models.ErrorCodeRegistrationDBError,
			Description: err.Error(),
		}
	}

	session := models.Session{
		ID:       uuid.New().String(),
		UserID:   createdUser.ID,
		UserRole: createdUser.Role,
		Token:    uuid.New().String(),
	}

	createdSession, err := r.repository.YDB.Sessions.Create(ctx, session)

	if err != nil {
		sentry.CaptureException(err)
		return nil, &models.Error{
			Code:        models.ErrorCodeRegistrationCreateSession,
			Description: err.Error(),
		}
	}

	return createdSession, nil
}
