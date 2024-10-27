package sessions

import (
	"context"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/constants"
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

	err := s.repository.YDB.Sessions.DeleteByID(ctx, id)

	if errors.Is(err, constants.ErrDBEntityNotFound) {
		return &models.Error{
			Code:        models.ErrorCodeSessionsDeleteSessionNotFound,
			Description: "Session with the given id was not found",
			StatusCode:  http.StatusBadRequest,
		}
	}

	if err != nil {
		sentry.CaptureException(err)
		return &models.Error{
			Code:        models.ErrorCodeSessionsDeleteSessionDBError,
			Description: err.Error(),
		}
	}

	return nil
}
