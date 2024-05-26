package ydbstarter

import (
	"os"

	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository/ydb-starter/registrations"
	ydb "github.com/ydb-platform/gorm-driver"
	environ "github.com/ydb-platform/ydb-go-sdk-auth-environ"
	"gorm.io/gorm"
)

type YdbStarter struct {
	Registrations *registrations.Registrations
	db            *gorm.DB
	config        *config.YdbStarter
}

func New(
	logger logger.Logger,
	config *config.YdbStarter,
) (*YdbStarter, error) {
	db := &gorm.DB{}

	return &YdbStarter{
		Registrations: registrations.New(db, logger),
		db:            db,
		config:        config,
	}, nil
}

func (y *YdbStarter) Connect() error {
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

func (y *YdbStarter) Disconnect() error {
	if y.db == nil {
		return nil
	}

	db, err := y.db.DB()

	if err != nil {
		return err
	}

	return db.Close()
}

func (y *YdbStarter) AutoMigrate() error {
	return y.db.AutoMigrate(
		&registrations.Registration{},
	)
}
