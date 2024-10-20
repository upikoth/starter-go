package ydb

import (
	"os"

	"github.com/pkg/errors"
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	passwordrecoveryrequests "github.com/upikoth/starter-go/internal/repository/ydb/password-recovery-requests"
	"github.com/upikoth/starter-go/internal/repository/ydb/registrations"
	"github.com/upikoth/starter-go/internal/repository/ydb/sessions"
	"github.com/upikoth/starter-go/internal/repository/ydb/users"
	ydbmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydb-models"
	ydb "github.com/ydb-platform/gorm-driver"
	environ "github.com/ydb-platform/ydb-go-sdk-auth-environ"
	"gorm.io/gorm"
)

type YDB struct {
	Users                    *users.Users
	Sessions                 *sessions.Sessions
	Registrations            *registrations.Registrations
	PasswordRecoveryRequests *passwordrecoveryrequests.PasswordRecoveryRequests
	db                       *gorm.DB
	config                   *config.Ydb
}

func New(
	logger logger.Logger,
	config *config.Ydb,
) (*YDB, error) {
	db := &gorm.DB{}

	return &YDB{
		Users:                    users.New(db, logger),
		Sessions:                 sessions.New(db, logger),
		Registrations:            registrations.New(db, logger),
		PasswordRecoveryRequests: passwordrecoveryrequests.New(db, logger),
		db:                       db,
		config:                   config,
	}, nil
}

func (y *YDB) Connect() error {
	filePath := y.config.AuthFileDirName + "/" + y.config.AuthFileName

	if len(y.config.YcSaJSONCredentials) > 0 {
		_, err := os.Stat(y.config.AuthFileDirName)

		if err != nil {
			mkdirErr := os.Mkdir(y.config.AuthFileDirName, 0777)

			if mkdirErr != nil {
				return errors.WithStack(mkdirErr)
			}
		}

		err = os.WriteFile(filePath, y.config.YcSaJSONCredentials, 0600)

		if err != nil {
			return errors.WithStack(err)
		}
	}

	err := os.Setenv("YDB_SERVICE_ACCOUNT_KEY_FILE_CREDENTIALS", filePath)

	if err != nil {
		return errors.WithStack(err)
	}

	db, err := gorm.Open(
		ydb.Open(y.config.Dsn, ydb.With(environ.WithEnvironCredentials())),
	)

	if err != nil {
		return errors.WithStack(err)
	}

	*y.db = *db

	return y.AutoMigrate()
}

func (y *YDB) Disconnect() error {
	if y.db == nil {
		return nil
	}

	db, err := y.db.DB()

	if err != nil {
		return errors.WithStack(err)
	}

	err = db.Close()

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (y *YDB) AutoMigrate() error {
	err := y.db.AutoMigrate(
		&ydbmodels.Registration{},
		&ydbmodels.User{},
		&ydbmodels.Session{},
		&ydbmodels.PasswordRecoveryRequest{},
	)

	return errors.WithStack(err)
}

func (y *YDB) Transaction(fn func(ydb *YDB) error) (err error) {
	return y.db.Transaction(func(tx *gorm.DB) error {
		return fn(y.WithTx(tx))
	})
}

func (y *YDB) WithTx(tx *gorm.DB) *YDB {
	return &YDB{
		Users:                    y.Users.WithTx(tx),
		Sessions:                 y.Sessions.WithTx(tx),
		Registrations:            y.Registrations.WithTx(tx),
		PasswordRecoveryRequests: y.PasswordRecoveryRequests.WithTx(tx),
		db:                       tx,
		config:                   y.config,
	}
}
