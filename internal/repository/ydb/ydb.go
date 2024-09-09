package ydb

import (
	"os"

	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	passwordrecoveryrequests "github.com/upikoth/starter-go/internal/repository/ydb/password-recovery-requests"
	passwordrecoveryrequestsandusers "github.com/upikoth/starter-go/internal/repository/ydb/password-recovery-requests-and-users"
	"github.com/upikoth/starter-go/internal/repository/ydb/registrations"
	registrationsandusers "github.com/upikoth/starter-go/internal/repository/ydb/registrations-and-users"
	"github.com/upikoth/starter-go/internal/repository/ydb/sessions"
	"github.com/upikoth/starter-go/internal/repository/ydb/users"
	ydbsmodels "github.com/upikoth/starter-go/internal/repository/ydb/ydbs-models"
	ydb "github.com/ydb-platform/gorm-driver"
	environ "github.com/ydb-platform/ydb-go-sdk-auth-environ"
	"gorm.io/gorm"
)

type Ydb struct {
	Registrations                    *registrations.Registrations
	RegistrationsAndUsers            *registrationsandusers.RegistrationsAndUsers
	Sessions                         *sessions.Sessions
	Users                            *users.Users
	PasswordRecoveryRequests         *passwordrecoveryrequests.PasswordRecoveryRequests
	PasswordRecoveryRequestsAndUsers *passwordrecoveryrequestsandusers.PasswordRecoveryRequestsAndUsers
	db                               *gorm.DB
	config                           *config.Ydb
}

func New(
	logger logger.Logger,
	config *config.Ydb,
) (*Ydb, error) {
	db := &gorm.DB{}

	return &Ydb{
		Registrations:                    registrations.New(db, logger),
		RegistrationsAndUsers:            registrationsandusers.New(db, logger),
		Sessions:                         sessions.New(db, logger),
		Users:                            users.New(db, logger),
		PasswordRecoveryRequests:         passwordrecoveryrequests.New(db, logger),
		PasswordRecoveryRequestsAndUsers: passwordrecoveryrequestsandusers.New(db, logger),
		db:                               db,
		config:                           config,
	}, nil
}

func (y *Ydb) Connect() error {
	filePath := y.config.AuthFileDirName + "/" + y.config.AuthFileName

	if len(y.config.YcSaJSONCredentials) > 0 {
		_, err := os.Stat(y.config.AuthFileDirName)

		if err != nil {
			mkdirErr := os.Mkdir(y.config.AuthFileDirName, 0777)

			if mkdirErr != nil {
				return mkdirErr
			}
		}

		err = os.WriteFile(filePath, y.config.YcSaJSONCredentials, 0600)

		if err != nil {
			return err
		}
	}

	os.Setenv("YDB_SERVICE_ACCOUNT_KEY_FILE_CREDENTIALS", filePath)

	db, err := gorm.Open(
		ydb.Open(y.config.Dsn, ydb.With(environ.WithEnvironCredentials())),
	)

	if err != nil {
		return err
	}

	*y.db = *db

	err = y.AutoMigrate()
	return err
}

func (y *Ydb) Disconnect() error {
	if y.db == nil {
		return nil
	}

	db, err := y.db.DB()

	if err != nil {
		return err
	}

	return db.Close()
}

func (y *Ydb) AutoMigrate() error {
	return y.db.AutoMigrate(
		&ydbsmodels.Registration{},
		&ydbsmodels.User{},
		&ydbsmodels.Session{},
		&ydbsmodels.PasswordRecoveryRequest{},
	)
}