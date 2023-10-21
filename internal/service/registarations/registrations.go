package registrations

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
	"github.com/upikoth/starter-go/internal/constants"
	"github.com/upikoth/starter-go/internal/model"
	emailSender "github.com/upikoth/starter-go/internal/pkg/email-sender"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository"
	usersService "github.com/upikoth/starter-go/internal/service/users"
	"golang.org/x/crypto/bcrypt"
)

type Registrations struct {
	logger     logger.Logger
	repository *repository.Repository

	usersService *usersService.Users

	config *config
}

type config struct {
	SiteURL                      string `envconfig:"SITE_URL" required:"true"`
	RegistrationConfirmationPath string `envconfig:"REGISTRATION_CONFIRMATION_PATH" required:"true"`
}

func newConfig() (*config, error) {
	config := &config{}

	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}

	return config, nil
}

func New(
	logger logger.Logger,
	repository *repository.Repository,
	usersService *usersService.Users,
) *Registrations {
	config, configErr := newConfig()

	if configErr != nil {
		log.Fatal(configErr)
	}

	return &Registrations{
		logger,
		repository,
		usersService,
		config,
	}
}

func (r *Registrations) Create(name string, email string, password string) *model.ExtendedError {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return &model.ExtendedError{
			ResponseCode:         http.StatusInternalServerError,
			ResponseErrorCode:    constants.ErrUnkownError,
			ResponseErrorDetails: err,
		}
	}

	registration := model.Registration{
		Name:                          name,
		Email:                         email,
		PasswordHash:                  string(hashedBytes),
		RegistrationConfirmationToken: uuid.New().String(),
	}

	user, _ := r.usersService.GetByEmail(registration.Email)

	if user.ID > 0 {
		return &model.ExtendedError{
			ResponseCode:      http.StatusBadRequest,
			ResponseErrorCode: constants.ErrRegistrationEmailAlreadyExist,
		}
	}

	emailSender, err := emailSender.New()

	if err != nil {
		return &model.ExtendedError{
			ResponseCode:         http.StatusInternalServerError,
			ResponseErrorCode:    constants.ErrUnkownError,
			ResponseErrorDetails: err,
		}
	}

	err = emailSender.Send(
		email,
		"Подтверждение регистрации",
		fmt.Sprintf(
			"Для завершения регистрации перейдите по ссылке: %s%s?token=%s",
			r.config.SiteURL,
			r.config.RegistrationConfirmationPath,
			registration.RegistrationConfirmationToken,
		),
	)

	if err != nil {
		return &model.ExtendedError{
			ResponseCode:         http.StatusInternalServerError,
			ResponseErrorCode:    constants.ErrUnkownError,
			ResponseErrorDetails: err,
		}
	}

	err = r.repository.Registrations.Create(registration)

	if err != nil {
		return &model.ExtendedError{
			ResponseCode:         http.StatusInternalServerError,
			ResponseErrorCode:    constants.ErrDbError,
			ResponseErrorDetails: err,
		}
	}

	return nil
}

func (r *Registrations) Confirm(token string) *model.ExtendedError {
	registration, err := r.repository.Registrations.GetByToken(token)

	if err != nil {
		return &model.ExtendedError{
			ResponseCode:         http.StatusInternalServerError,
			ResponseErrorCode:    constants.ErrDbError,
			ResponseErrorDetails: err,
		}
	}

	user := model.User{
		Name:         registration.Name,
		Email:        registration.Email,
		Status:       model.UserStatusActive,
		PasswordHash: registration.PasswordHash,
	}

	err = r.usersService.Create(user)

	if err != nil {
		return &model.ExtendedError{
			ResponseCode:         http.StatusInternalServerError,
			ResponseErrorCode:    constants.ErrDbError,
			ResponseErrorDetails: err,
		}
	}

	err = r.repository.Registrations.DeleteByID(registration.ID)

	if err != nil {
		return &model.ExtendedError{
			ResponseCode:         http.StatusInternalServerError,
			ResponseErrorCode:    constants.ErrDbError,
			ResponseErrorDetails: err,
		}
	}

	return nil
}
