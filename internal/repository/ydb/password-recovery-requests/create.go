package passwordrecoveryrequests

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/models"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/query"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func (p *PasswordRecoveryRequests) Create(
	inputCtx context.Context,
	prrToCreate *models.PasswordRecoveryRequest,
) (res *models.PasswordRecoveryRequest, err error) {
	tracer := otel.Tracer("Repository: YDB.PasswordRecoveryRequests.Create")
	ctx, span := tracer.Start(inputCtx, "Repository: YDB.PasswordRecoveryRequests.Create")
	defer span.End()

	defer func() {
		if err != nil {
			span.RecordError(err)
		} else {
			bytes, _ := json.Marshal(res)
			span.SetAttributes(
				attribute.String("ydb.res", string(bytes)),
			)
		}
	}()

	var dbCreatedPRR ydbmodels.PasswordRecoveryRequest
	dbPRRToCreate := ydbmodels.NewYDBPasswordRecoveryRequestModel(prrToCreate)

	err = p.executeInQueryTransaction(ctx, func(qCtx context.Context, tx query.Transaction) error {
		qRes, qErr := tx.QueryResultSet(
			qCtx,
			`declare $id as text;
				declare $email as text;
				declare $confirmation_token as text;
				
				insert into password_recovery_requests
				(id, email, confirmation_token)
				values ($id, $email, $confirmation_token);

				select
					id,
					email,
					confirmation_token,
				from password_recovery_requests as prr
				where prr.id = $id;`,
			query.WithParameters(
				ydb.ParamsBuilder().
					Param("$id").Text(dbPRRToCreate.ID).
					Param("$email").Text(dbPRRToCreate.Email).
					Param("$confirmation_token").Text(dbPRRToCreate.ConfirmationToken).
					Build(),
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
				query.Named("id", &dbCreatedPRR.ID),
				query.Named("email", &dbCreatedPRR.Email),
				query.Named("confirmation_token", &dbCreatedPRR.ConfirmationToken),
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

	return dbCreatedPRR.FromYDBModel(), nil
}
