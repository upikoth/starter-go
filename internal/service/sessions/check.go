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
) (models.Session, error) {
	span := sentry.StartSpan(inputCtx, "Service: Sessions.CheckSessionToken")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	session, err := s.repository.YDB.Sessions.GetByToken(ctx, token)

	if err != nil {
		sentry.CaptureException(err)
		return session, &models.Error{
			Code:        models.ErrorCodeUserUnauthorized,
			Description: err.Error(),
			StatusCode:  http.StatusUnauthorized,
		}
	}

	if session.ID == "" {
		return session, &models.Error{
			Code:        models.ErrorCodeUserUnauthorized,
			Description: "User session is invalid",
			StatusCode:  http.StatusUnauthorized,
		}
	}

	return session, err
}
