package ydb

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/pressly/goose/v3"
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	passwordrecoveryrequests "github.com/upikoth/starter-go/internal/repository/ydb/password-recovery-requests"
	"github.com/upikoth/starter-go/internal/repository/ydb/registrations"
	"github.com/upikoth/starter-go/internal/repository/ydb/sessions"
	"github.com/upikoth/starter-go/internal/repository/ydb/users"
	ydbOtel "github.com/ydb-platform/ydb-go-sdk-otel"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/query"
	"github.com/ydb-platform/ydb-go-sdk/v3/trace"
	yc "github.com/ydb-platform/ydb-go-yc"
)

type YDB struct {
	Users                    *users.Users
	Sessions                 *sessions.Sessions
	Registrations            *registrations.Registrations
	PasswordRecoveryRequests *passwordrecoveryrequests.PasswordRecoveryRequests
	db                       *ydb.Driver
	config                   *config.Ydb
	logger                   logger.Logger
}

func New(
	logger logger.Logger,
	config *config.Ydb,
) (*YDB, error) {
	return &YDB{
		config: config,
		logger: logger,
	}, nil
}

func (y *YDB) Connect(ctx context.Context) error {
	filePath := fmt.Sprintf("%s/%s", y.config.AuthFileDirName, y.config.AuthFileName)
	err := writeCredentialsToFile(y.config.AuthFileDirName, y.config.AuthFileName, y.config.YcSaJSONCredentials)

	if err != nil {
		return err
	}

	driver, err := ydb.Open(
		ctx,
		y.config.Dsn,
		yc.WithServiceAccountKeyFileCredentials(filePath),
		ydbOtel.WithTraces(
			ydbOtel.WithDetails(trace.DetailsAll),
		),
	)

	if err != nil {
		return errors.WithStack(err)
	}

	err = y.Migrate(ctx, driver)

	if err != nil {
		return errors.WithStack(err)
	}

	y.db = driver
	y.Users = users.New(y.db, y.logger)
	y.Sessions = sessions.New(y.db, y.logger)
	y.Registrations = registrations.New(y.db, y.logger)
	y.PasswordRecoveryRequests = passwordrecoveryrequests.New(y.db, y.logger)

	return nil
}

//go:embed 1_migrations/*.sql
var embedMigrations embed.FS

func (y *YDB) Migrate(ctx context.Context, driver *ydb.Driver) error {
	connector, cErr := ydb.Connector(
		driver,
		ydb.WithDefaultQueryMode(ydb.ScriptingQueryMode),
		ydb.WithFakeTx(ydb.ScriptingQueryMode),
		ydb.WithAutoDeclare(),
		ydb.WithNumericArgs(),
	)
	if cErr != nil {
		return errors.WithStack(cErr)
	}

	sqlDB := sql.OpenDB(connector)

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("ydb"); err != nil {
		return errors.WithStack(err)
	}

	if err := goose.UpContext(ctx, sqlDB, "1_migrations"); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (y *YDB) Disconnect(ctx context.Context) error {
	if y.db == nil {
		return nil
	}

	return errors.WithStack(y.db.Close(ctx))
}

func writeCredentialsToFile(dirName string, fileName string, credentials []byte) error {
	filePath := fmt.Sprintf("%s/%s", dirName, fileName)

	if len(credentials) > 0 {
		_, err := os.Stat(dirName)

		if err != nil {
			mkdirErr := os.Mkdir(dirName, 0777)

			if mkdirErr != nil {
				return errors.WithStack(mkdirErr)
			}
		}

		err = os.WriteFile(filePath, credentials, 0600)

		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func (y *YDB) Transaction(
	ctx context.Context,
	fn func(ctxTx context.Context, ydb *YDB) error,
	opts ...query.TransactionOption,
) error {
	return y.db.Query().Do(ctx, func(ctx context.Context, s query.Session) error {
		tx, err := s.Begin(ctx, query.TxSettings(opts...))

		if err != nil {
			return errors.WithStack(err)
		}

		defer func() { _ = tx.Rollback(ctx) }()

		yTx := y.withTx(tx)
		if fnErr := fn(ctx, yTx); fnErr != nil {
			return fnErr
		}

		return tx.CommitTx(ctx)
	})
}

func (y *YDB) TransactionWithSerializeLevel() query.TransactionOption {
	return query.WithSerializableReadWrite()
}

func (y *YDB) withTx(tx query.Transaction) *YDB {
	return &YDB{
		Users:                    y.Users.WithTx(tx),
		Sessions:                 y.Sessions.WithTx(tx),
		Registrations:            y.Registrations.WithTx(tx),
		PasswordRecoveryRequests: y.PasswordRecoveryRequests.WithTx(tx),
		db:                       y.db,
		logger:                   y.logger,
		config:                   y.config,
	}
}
