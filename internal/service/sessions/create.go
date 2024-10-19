package sessions

import (
	"context"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	"github.com/upikoth/starter-go/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (s *Sessions) Create(
	inputCtx context.Context,
	params models.SessionCreateParams,
) (*models.Session, error) {
	span := sentry.StartSpan(inputCtx, "Service: Sessions.Create")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	user, err := s.repository.YDB.Users.GetByEmail(ctx, params.Email)

	if err != nil {
		sentry.CaptureException(err)
		return nil, &models.Error{
			Code:        models.ErrorCodeSessionsCreateSessionDBError,
			Description: err.Error(),
		}
	}

	if user.ID == "" {
		return nil, &models.Error{
			Code:        models.ErrorCodeSessionsCreateSessionWrongEmailOrPassword,
			Description: "Incorrect email or password",
			StatusCode:  http.StatusBadRequest,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(params.Password))

	if err != nil {
		return &models.Session{}, &models.Error{
			Code:        models.ErrorCodeSessionsCreateSessionWrongEmailOrPassword,
			Description: "Incorrect email or password",
			StatusCode:  http.StatusBadRequest,
		}
	}

	sessionToCreate := models.Session{
		ID:       uuid.New().String(),
		UserID:   user.ID,
		UserRole: user.Role,
		Token:    uuid.New().String(),
	}

	session, err := s.repository.YDB.Sessions.Create(ctx, sessionToCreate)

	if err != nil {
		sentry.CaptureException(err)
		return nil, &models.Error{
			Code:        models.ErrorCodeSessionsCreateSessionDBError,
			Description: err.Error(),
		}
	}

	return session, err
}
