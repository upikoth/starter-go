package config

import (
	"github.com/kelseyhightower/envconfig"
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
	YcpStarter YcpStarter
	YdbStarter YdbStarter
}

type YcpStarter struct {
	Host        string `envconfig:"YCP_STARTER_HOST" required:"true"`
	Port        string `envconfig:"YCP_STARTER_PORT" required:"true"`
	FromName    string `envconfig:"YCP_STARTER_FROM_NAME" required:"true"`
	FromAddress string `envconfig:"YCP_STARTER_FROM_ADDRESS" required:"true"`
	Username    string `envconfig:"YCP_STARTER_USERNAME" required:"true"`
	Password    string `envconfig:"YCP_STARTER_PASSWORD" required:"true"`
}

type YdbStarter struct {
	Dsn                 string `envconfig:"YDB_STARTER_DSN" required:"true"`
	AuthFileDirName     string `envconfig:"YDB_STARTER_AUTH_FILE_DIR_NAME" required:"true"`
	AuthFileName        string `envconfig:"YDB_STARTER_AUTH_FILE_NAME" required:"true"`
	YcSaJSONCredentials []byte `envconfig:"YC_SA_JSON_CREDENTIALS"`
}

func New() (*Config, error) {
	config := &Config{}

	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}

	return config, nil
}
