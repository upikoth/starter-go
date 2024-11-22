package sessions

import (
	"context"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/constants"
	"github.com/upikoth/starter-go/internal/models"
	"github.com/upikoth/starter-go/internal/pkg/tracing"
	"go.opentelemetry.io/otel"
)

func (s *Sessions) DeleteByID(
	inputCtx context.Context,
	id string,
) error {
	tracer := otel.Tracer(tracing.GetServiceTraceName())
	ctx, span := tracer.Start(inputCtx, tracing.GetServiceTraceName())
	defer span.End()

	err := s.repository.YDB.Sessions.DeleteByID(ctx, id)

	if errors.Is(err, constants.ErrDBEntityNotFound) {
		return &models.Error{
			Code:        models.ErrorCodeSessionsDeleteSessionNotFound,
			Description: "Session with the given id was not found",
			StatusCode:  http.StatusBadRequest,
		}
	}

	if err != nil {
		span.RecordError(err)
		sentry.CaptureException(err)
		return &models.Error{
			Code:        models.ErrorCodeSessionsDeleteSessionDBError,
			Description: err.Error(),
		}
	}

	return nil
}
