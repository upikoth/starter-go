package registrations

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/upikoth/starter-go/internal/constants"
	"github.com/upikoth/starter-go/internal/model"
	emailSender "github.com/upikoth/starter-go/internal/pkg/email-sender"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type Registrations struct {
	logger     logger.Logger
	repository *repository.Repository
}

func New(logger logger.Logger, repository *repository.Repository) *Registrations {
	return &Registrations{
		logger,
		repository,
	}
}

func (r *Registrations) Create(email string, password string) *model.ExtendedError {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return &model.ExtendedError{
			ResponseCode:         http.StatusInternalServerError,
			ResponseErrorCode:    constants.ErrUnkownError,
			ResponseErrorDetails: err,
		}
	}

	registration := model.Registration{
		Email:                         email,
		PasswordHash:                  string(hashedBytes),
		RegistrationConfirmationToken: uuid.New().String(),
	}

	user, _ := r.repository.Users.GetByEmail(registration.Email)

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
			"Для завершения регистрации перейдите по ссылке: https://starter.upikoth.ru?token=%s",
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
