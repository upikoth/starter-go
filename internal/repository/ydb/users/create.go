package users

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

func (u *Users) Create(
	inputCtx context.Context,
	userToCreate *models.User,
) (res *models.User, err error) {
	tracer := otel.Tracer("Repository: YDB.Users.Create")
	ctx, span := tracer.Start(inputCtx, "Repository: YDB.Users.Create")
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

	var dbCreatedUser ydbmodels.User
	dbUserToCreate := ydbmodels.NewYDBUserModel(userToCreate)

	err = u.executeInQueryTransaction(ctx, func(qCtx context.Context, tx query.Transaction) error {
		qRes, qErr := tx.QueryResultSet(
			qCtx,
			`declare $id as text;
				declare $email as text;
				declare $role as text;
				declare $password_hash as text;
	
				insert into users
				(id, email, role, password_hash)
				values ($id, $email, $role, $password_hash);
	
				select
					id,
					email,
					role,
					password_hash,
				from users
				where users.id = $id;`,
			query.WithParameters(
				ydb.ParamsBuilder().
					Param("$id").Text(dbUserToCreate.ID).
					Param("$email").Text(dbUserToCreate.Email).
					Param("$role").Text(dbUserToCreate.Role).
					Param("$password_hash").Text(dbUserToCreate.PasswordHash).
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
				query.Named("id", &dbCreatedUser.ID),
				query.Named("email", &dbCreatedUser.Email),
				query.Named("role", &dbCreatedUser.Role),
				query.Named("password_hash", &dbCreatedUser.PasswordHash),
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

	return dbCreatedUser.FromYDBModel(), nil
}
