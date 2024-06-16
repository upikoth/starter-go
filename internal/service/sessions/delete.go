package sessions

import (
	"context"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/upikoth/starter-go/internal/models"
)

func (s *Sessions) DeleteByID(
	inputCtx context.Context,
	id string,
) error {
	span := sentry.StartSpan(inputCtx, "Service: Sessions.DeleteByID")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	session, err := s.repository.YdbStarter.Sessions.GetByID(ctx, id)

	if err != nil {
		sentry.CaptureException(err)
		return &models.Error{
			Code:        models.ErrorCodeSessionsDeleteSessionDbError,
			Description: err.Error(),
		}
	}

	if session.ID == "" {
		return &models.Error{
			Code:        models.ErrorCodeSessionsDeleteSessionNotFound,
			Description: "Session with the given id was not found",
			StatusCode:  http.StatusBadRequest,
		}
	}

	err = s.repository.YdbStarter.Sessions.DeleteByID(ctx, id)

	if err != nil {
		sentry.CaptureException(err)
		return &models.Error{
			Code:        models.ErrorCodeSessionsDeleteSessionDbError,
			Description: err.Error(),
		}
	}

	return err
}
