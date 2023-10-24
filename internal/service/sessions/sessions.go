package sessions

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/upikoth/starter-go/internal/constants"
	"github.com/upikoth/starter-go/internal/model"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository"
	usersService "github.com/upikoth/starter-go/internal/service/users"
	"golang.org/x/crypto/bcrypt"
)

type Sessions struct {
	logger     logger.Logger
	repository *repository.Repository

	usersService *usersService.Users

	config *config
}

type config struct {
	SessionLifetimeLengthDays int `envconfig:"SESSION_LIFETIME_LENGTH_DAYS" required:"true"`
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
) *Sessions {
	config, configErr := newConfig()

	if configErr != nil {
		log.Fatal(configErr)
	}

	return &Sessions{
		logger,
		repository,
		usersService,
		config,
	}
}

func (s *Sessions) GetAll() ([]model.Session, error) {
	return s.repository.Sessions.GetAll()
}

func (s *Sessions) Create(email string, password string, userAgent string) (model.Session, *model.ExtendedError) {
	user, _ := s.usersService.GetByEmail(email)

	if user.ID == 0 {
		return model.Session{}, &model.ExtendedError{
			ResponseCode:      http.StatusBadRequest,
			ResponseErrorCode: constants.ErrSessionCreateUserOrPasswordInvalid,
		}
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) != nil {
		return model.Session{}, &model.ExtendedError{
			ResponseCode:      http.StatusBadRequest,
			ResponseErrorCode: constants.ErrSessionCreateUserOrPasswordInvalid,
		}
	}

	createdAt := time.Now().Local().UTC()
	expiredAt := createdAt.Add(time.Hour * 24 * time.Duration(s.config.SessionLifetimeLengthDays))

	tokenByteLength := 16
	token := make([]byte, tokenByteLength)
	_, err := rand.Read(token)

	if err != nil {
		return model.Session{}, &model.ExtendedError{
			ResponseCode:         http.StatusBadRequest,
			ResponseErrorCode:    constants.ErrUnkownError,
			ResponseErrorDetails: err,
		}
	}

	session := model.Session{
		Token:     fmt.Sprintf("%x", token),
		UserID:    user.ID,
		UserAgent: userAgent,
		CreatedAt: createdAt.Format(time.RFC3339Nano),
		ExpiredAt: expiredAt.Format(time.RFC3339Nano),
	}

	err = s.repository.Sessions.Create(session)

	if err != nil {
		return model.Session{}, &model.ExtendedError{
			ResponseCode:         http.StatusBadRequest,
			ResponseErrorCode:    constants.ErrDbError,
			ResponseErrorDetails: err,
		}
	}

	return session, nil
}
