package sessions

import (
	"context"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/constants"
	"github.com/upikoth/starter-go/internal/models"
)

func (s *Sessions) CheckToken(
	inputCtx context.Context,
	token string,
) (*models.SessionWithUserRole, error) {
	span := sentry.StartSpan(inputCtx, "Service: Sessions.CheckSessionToken")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	session, err := s.repository.YDB.Sessions.GetByToken(ctx, token)

	if errors.Is(err, constants.ErrDBEntityNotFound) {
		return nil, &models.Error{
			Code:        models.ErrorCodeUserUnauthorized,
			Description: "User session is invalid",
			StatusCode:  http.StatusUnauthorized,
		}
	}

	if err != nil {
		sentry.CaptureException(err)
		return nil, &models.Error{
			Code:        models.ErrorCodeSessionsCheckTokenDBError,
			Description: err.Error(),
		}
	}

	return session, err
}
