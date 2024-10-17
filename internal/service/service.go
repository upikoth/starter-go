package service

import (
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository"
	passwordrecoveryrequests "github.com/upikoth/starter-go/internal/service/password-recovery-requests"
	"github.com/upikoth/starter-go/internal/service/registrations"
	"github.com/upikoth/starter-go/internal/service/sessions"
	"github.com/upikoth/starter-go/internal/service/users"
)

type Service struct {
	Registrations            *registrations.Registrations
	Sessions                 *sessions.Sessions
	PasswordRecoveryRequests *passwordrecoveryrequests.PasswordRecoveryRequests
	Users                    *users.Users
}

func New(
	logger logger.Logger,
	config *config.Config,
	repository *repository.Repository,
) (*Service, error) {
	return &Service{
		Registrations: registrations.New(
			logger,
			&config.Service.Registrations,
			repository,
		),
		Sessions: sessions.New(
			logger,
			repository,
		),
		PasswordRecoveryRequests: passwordrecoveryrequests.New(
			logger,
			&config.Service.PasswordRecoveryRequests,
			repository,
		),
		Users: users.New(
			logger,
			repository,
		),
	}, nil
}
