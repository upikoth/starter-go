package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Config struct {
	Environment string `envconfig:"ENVIRONMENT" required:"true"`
	Controller  Controller
	Service     Service
	Repository  Repository
}

type Controller struct {
	HTTP ControllerHTTP
}

type ControllerHTTP struct {
	Port        string `envconfig:"PORT" required:"true"`
	SentryDsn   string `envconfig:"SENTRY_DSN" required:"true"`
	Environment string `envconfig:"ENVIRONMENT" required:"true"`
}
type Service struct {
	Registrations            Registrations
	PasswordRecoveryRequests PasswordRecoveryRequests
}

type Registrations struct {
	FrontURL                         string `envconfig:"FRONT_URL" required:"true"`
	FrontConfirmationRegistrationURL string `envconfig:"FRONT_CONFIRMATION_REGISTRATION_URL" required:"true"`
}

type PasswordRecoveryRequests struct {
	FrontURL string `envconfig:"FRONT_URL" required:"true"`
	//nolint:lll
	FrontConfirmationPasswordRecoveryRequestURL string `envconfig:"FRONT_CONFIRMATION_PASSWORD_RECOVERY_REQUEST_URL" required:"true"`
}

type Repository struct {
	Ycp Ycp
	Ydb Ydb
}

type Ycp struct {
	Host        string `envconfig:"YCP_HOST" required:"true"`
	Port        string `envconfig:"YCP_PORT" required:"true"`
	FromName    string `envconfig:"YCP_FROM_NAME" required:"true"`
	FromAddress string `envconfig:"YCP_FROM_ADDRESS" required:"true"`
	Username    string `envconfig:"YCP_USERNAME" required:"true"`
	Password    string `envconfig:"YCP_PASSWORD" required:"true"`
}

type Ydb struct {
	Dsn                 string `envconfig:"YDB_DSN" required:"true"`
	AuthFileDirName     string `envconfig:"YDB_AUTH_FILE_DIR_NAME" required:"true"`
	AuthFileName        string `envconfig:"YDB_AUTH_FILE_NAME" required:"true"`
	YcSaJSONCredentials []byte `envconfig:"YC_SA_JSON_CREDENTIALS"`
}

func New() (*Config, error) {
	config := &Config{}

	if err := envconfig.Process("", config); err != nil {
		return nil, errors.WithStack(err)
	}

	return config, nil
}
