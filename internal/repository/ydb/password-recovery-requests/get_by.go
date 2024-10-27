package passwordrecoveryrequests

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/constants"
	"github.com/upikoth/starter-go/internal/models"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/query"
)

type fieldNameGetBy string

var (
	fieldNameGetByConfrimationToken fieldNameGetBy = "confirmation_token"
	fieldNameGetByEmail             fieldNameGetBy = "email"
)

func (p *PasswordRecoveryRequests) getBy(
	inputCtx context.Context,
	fieldName fieldNameGetBy,
	fieldValue string,
) (res *models.PasswordRecoveryRequest, err error) {
	span := sentry.StartSpan(inputCtx, "Repository: YDB.PasswordRecoveryRequests.getBy")
	defer func() {
		if err != nil && !errors.Is(err, constants.ErrDBEntityNotFound) {
			sentry.CaptureException(err)
		} else {
			bytes, _ := json.Marshal(res)
			span.SetData("Result", string(bytes))
		}
		span.Finish()
	}()
	ctx := span.Context()

	var prr ydbmodels.PasswordRecoveryRequest

	err = p.executeInQueryTransaction(ctx, func(qCtx context.Context, tx query.Transaction) error {
		qRes, qErr := tx.QueryResultSet(
			qCtx,
			fmt.Sprintf(
				`declare $filterValue as text;
					select
						id,
						email,
						confirmation_token,
					from password_recovery_requests
					where %s = $filterValue;`,
				fieldName,
			),
			query.WithParameters(
				ydb.ParamsBuilder().Param("$filterValue").Text(fieldValue).Build(),
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
				query.Named("id", &prr.ID),
				query.Named("email", &prr.Email),
				query.Named("confirmation_token", &prr.ConfirmationToken),
			)

			if sErr != nil {
				return errors.WithStack(sErr)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	if reflect.ValueOf(prr).IsZero() {
		return nil, errors.WithStack(constants.ErrDBEntityNotFound)
	}

	return prr.FromYDBModel(), nil
}
