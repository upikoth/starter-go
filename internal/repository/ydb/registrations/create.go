package registrations

import (
	"context"
	"encoding/json"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/models"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/query"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func (r *Registrations) Create(
	inputCtx context.Context,
	registrationToCreate *models.Registration,
) (res *models.Registration, err error) {
	tracer := otel.Tracer("Repository: YDB.Registrations.Create")
	ctx, span := tracer.Start(inputCtx, "Repository: YDB.Registrations.Create")
	defer span.End()

	defer func() {
		if err != nil {
			span.RecordError(err)
			sentry.CaptureException(err)
		} else {
			bytes, _ := json.Marshal(res)
			span.SetAttributes(
				attribute.String("ydb.res", string(bytes)),
			)
		}
	}()

	var dbCreatedRegistration ydbmodels.Registration
	dbRegistrationToCreate := ydbmodels.NewYDBRegistrationModel(registrationToCreate)

	err = r.executeInQueryTransaction(ctx, func(qCtx context.Context, tx query.Transaction) error {
		qRes, qErr := tx.QueryResultSet(
			qCtx,
			`declare $id as text;
				declare $email as text;
				declare $confirmation_token as text;
				
				insert into registrations
				(id, email, confirmation_token)
				values ($id, $email, $confirmation_token);

				select
					id,
					email,
					confirmation_token,
				from registrations
				where registrations.id = $id;`,
			query.WithParameters(
				ydb.ParamsBuilder().
					Param("$id").Text(dbRegistrationToCreate.ID).
					Param("$email").Text(dbRegistrationToCreate.Email).
					Param("$confirmation_token").Text(dbRegistrationToCreate.ConfirmationToken).
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
				query.Named("id", &dbCreatedRegistration.ID),
				query.Named("email", &dbCreatedRegistration.Email),
				query.Named("confirmation_token", &dbCreatedRegistration.ConfirmationToken),
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

	return dbCreatedRegistration.FromYDBModel(), nil
}
