package registrations

import (
	"context"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/constants"
	"github.com/upikoth/starter-go/internal/models"
	"github.com/upikoth/starter-go/internal/repository/ydb"
	"go.opentelemetry.io/otel"
	"golang.org/x/crypto/bcrypt"
)

func (r *Registrations) Confirm(
	inputCtx context.Context,
	params models.RegistrationConfirmParams,
) (*models.SessionWithUserRole, error) {
	tracer := otel.Tracer("Service: Registrations.Confirm")
	ctx, span := tracer.Start(inputCtx, "Service: Registrations.Confirm")
	defer span.End()

	registration, err := r.repository.YDB.Registrations.GetByToken(ctx, params.ConfirmationToken)

	if errors.Is(err, constants.ErrDBEntityNotFound) {
		return nil, &models.Error{
			Code:        models.ErrorCodeRegistrationRegistrationNotFound,
			Description: "Registration with transferred token not found",
			StatusCode:  http.StatusBadRequest,
		}
	}

	if err != nil {
		span.RecordError(err)
		sentry.CaptureException(err)
		return nil, &models.Error{
			Code:        models.ErrorCodeRegistrationYdbCheckConfirmationToken,
			Description: err.Error(),
		}
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)

	if err != nil {
		span.RecordError(err)
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
	err = r.repository.YDB.Transaction(
		ctx,
		func(ctxTx context.Context, ydbTx *ydb.YDB) error {
			dbErr := ydbTx.Registrations.DeleteByID(ctxTx, registration.ID)

			if dbErr != nil {
				return dbErr
			}

			createdUser, dbErr = ydbTx.Users.Create(ctxTx, &newUser)

			if dbErr != nil {
				return dbErr
			}

			return nil
		},
		r.repository.YDB.TransactionWithSerializeLevel(),
	)

	if err != nil {
		span.RecordError(err)
		sentry.CaptureException(err)
		return nil, &models.Error{
			Code:        models.ErrorCodeRegistrationDBError,
			Description: err.Error(),
		}
	}

	session := &models.Session{
		ID:     uuid.New().String(),
		Token:  uuid.New().String(),
		UserID: createdUser.ID,
	}

	createdSession, err := r.repository.YDB.Sessions.Create(ctx, session)

	if err != nil {
		span.RecordError(err)
		sentry.CaptureException(err)
		return nil, &models.Error{
			Code:        models.ErrorCodeRegistrationCreateSession,
			Description: err.Error(),
		}
	}

	return createdSession, nil
}
