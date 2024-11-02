package sessions

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/constants"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/query"
	"go.opentelemetry.io/otel"
)

func (s *Sessions) DeleteByID(
	inputCtx context.Context,
	id string,
) (err error) {
	tracer := otel.Tracer("Repository: YDB.Sessions.DeleteByID")
	ctx, span := tracer.Start(inputCtx, "Repository: YDB.Sessions.DeleteByID")
	defer span.End()

	defer func() {
		if err != nil {
			span.RecordError(err)
		}
	}()

	var deletedSession ydbmodels.Session

	err = s.executeInQueryTransaction(ctx, func(qCtx context.Context, tx query.Transaction) error {
		qRes, qErr := tx.QueryResultSet(
			qCtx,
			`declare $id as text;
				delete from sessions
				where id = $id
				returning id, token, user_id`,
			query.WithParameters(
				ydb.ParamsBuilder().Param("$id").Text(id).Build(),
			),
		)

		if qErr != nil {
			return errors.WithStack(qErr)
		}

		defer func() { _ = qRes.Close(qCtx) }()

		for row, rErr := range qRes.Rows(qCtx) {
			if rErr != nil {
				return errors.WithStack(rErr)
			}

			sErr := row.ScanNamed(
				query.Named("id", &deletedSession.ID),
			)

			if sErr != nil {
				return errors.WithStack(sErr)
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	if reflect.ValueOf(deletedSession).IsZero() {
		return errors.WithStack(constants.ErrDBEntityNotFound)
	}

	return nil
}
