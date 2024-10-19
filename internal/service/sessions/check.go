package sessions

import (
	"context"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/upikoth/starter-go/internal/models"
)

func (s *Sessions) CheckToken(
	inputCtx context.Context,
	token string,
) (*models.Session, error) {
	span := sentry.StartSpan(inputCtx, "Service: Sessions.CheckSessionToken")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	session, err := s.repository.YDB.Sessions.GetByToken(ctx, token)

	if err != nil {
		sentry.CaptureException(err)
		return nil, &models.Error{
			Code:        models.ErrorCodeSessionsCheckTokenDBError,
			Description: err.Error(),
		}
	}

	if session.ID == "" {
		return nil, &models.Error{
			Code:        models.ErrorCodeUserUnauthorized,
			Description: "User session is invalid",
			StatusCode:  http.StatusUnauthorized,
		}
	}

	return session, err
}
