package sessions

import (
	"context"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/constants"
	"github.com/upikoth/starter-go/internal/models"
	"go.opentelemetry.io/otel"
)

func (s *Sessions) CheckToken(
	inputCtx context.Context,
	token string,
) (*models.SessionWithUserRole, error) {
	tracer := otel.Tracer("Service: Sessions.CheckToken")
	ctx, span := tracer.Start(inputCtx, "Service: Sessions.CheckToken")
	defer span.End()

	session, err := s.repository.YDB.Sessions.GetByToken(ctx, token)

	if errors.Is(err, constants.ErrDBEntityNotFound) {
		return nil, &models.Error{
			Code:        models.ErrorCodeUserUnauthorized,
			Description: "User session is invalid",
			StatusCode:  http.StatusUnauthorized,
		}
	}

	if err != nil {
		span.RecordError(err)
		sentry.CaptureException(err)
		return nil, &models.Error{
			Code:        models.ErrorCodeSessionsCheckTokenDBError,
			Description: err.Error(),
		}
	}

	return session, err
}
