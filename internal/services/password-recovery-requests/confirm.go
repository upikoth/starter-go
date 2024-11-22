package passwordrecoveryrequests

import (
	"context"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/constants"
	"github.com/upikoth/starter-go/internal/models"
	"github.com/upikoth/starter-go/internal/pkg/tracing"
	"github.com/upikoth/starter-go/internal/repositories/ydb"
	"go.opentelemetry.io/otel"
	"golang.org/x/crypto/bcrypt"
)

func (p *PasswordRecoveryRequests) Confirm(
	inputCtx context.Context,
	params models.PasswordRecoveryRequestConfirmParams,
) (*models.SessionWithUserRole, error) {
	tracer := otel.Tracer(tracing.GetServiceTraceName())
	ctx, span := tracer.Start(inputCtx, tracing.GetServiceTraceName())
	defer span.End()

	passwordRecoveryRequest, err := p.
		repository.
		YDB.
		PasswordRecoveryRequests.
		GetByToken(ctx, params.ConfirmationToken)

	if errors.Is(err, constants.ErrDBEntityNotFound) {
		return nil, &models.Error{
			Code:        models.ErrorCodePasswordRecoveryRequestPasswordRecoveryRequestNotFound,
			Description: "Password recovery request with transferred token not found",
			StatusCode:  http.StatusBadRequest,
		}
	}

	if err != nil {
		span.RecordError(err)
		sentry.CaptureException(err)

		return nil, &models.Error{
			Code:        models.ErrorCodePasswordRecoveryRequestYdbCheckConfirmationToken,
			Description: err.Error(),
		}
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(params.NewPassword), bcrypt.DefaultCost)

	if err != nil {
		span.RecordError(err)
		sentry.CaptureException(err)

		return nil, &models.Error{
			Code:        models.ErrorCodePasswordRecoveryRequestGeneratePasswordHash,
			Description: err.Error(),
		}
	}

	var updatedUser *models.User
	err = p.repository.YDB.Transaction(
		ctx,
		func(ctxTx context.Context, ydbTx *ydb.YDB) error {
			user, dbErr := ydbTx.Users.GetByEmail(ctxTx, passwordRecoveryRequest.Email)

			if dbErr != nil {
				return dbErr
			}

			user.PasswordHash = string(passwordHash)

			dbUpdatedUser, dbErr := ydbTx.Users.Update(ctxTx, user)
			updatedUser = dbUpdatedUser

			if dbErr != nil {
				return dbErr
			}

			return ydbTx.PasswordRecoveryRequests.DeleteByID(ctxTx, passwordRecoveryRequest.ID)
		},
		p.repository.YDB.TransactionWithSerializeLevel(),
	)

	if err != nil {
		span.RecordError(err)
		sentry.CaptureException(err)

		return nil, &models.Error{
			Code:        models.ErrorCodePasswordRecoveryRequestUpdateUserPassword,
			Description: err.Error(),
		}
	}

	session := &models.Session{
		ID:     uuid.New().String(),
		Token:  uuid.New().String(),
		UserID: updatedUser.ID,
	}

	createdSession, err := p.repository.YDB.Sessions.Create(ctx, session)

	if err != nil {
		span.RecordError(err)
		sentry.CaptureException(err)

		return nil, &models.Error{
			Code:        models.ErrorCodePasswordRecoveryRequestCreateSession,
			Description: err.Error(),
		}
	}

	return createdSession, nil
}
