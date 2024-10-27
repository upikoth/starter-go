package passwordrecoveryrequests

import (
	"context"
	"reflect"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/constants"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/query"
)

func (p *PasswordRecoveryRequests) DeleteByID(
	inputCtx context.Context,
	id string,
) (err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.PasswordRecoveryRequests.DeleteByID")
	defer func() {
		if err != nil && !errors.Is(err, constants.ErrDBEntityNotFound) {
			sentry.CaptureException(err)
		}
		span.Finish()
	}()
	ctx := span.Context()

	var deletedPRR ydbmodels.PasswordRecoveryRequest

	err = p.executeInQueryTransaction(ctx, func(qCtx context.Context, tx query.Transaction) error {
		qRes, qErr := tx.QueryResultSet(
			qCtx,
			`declare $id as text;
				delete from password_recovery_requests
				where id = $id
				returning id, email, confirmation_token`,
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
				query.Named("id", &deletedPRR.ID),
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

	if reflect.ValueOf(deletedPRR).IsZero() {
		return errors.WithStack(constants.ErrDBEntityNotFound)
	}

	return nil
}
