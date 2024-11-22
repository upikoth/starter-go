package service

import (
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository"
	"github.com/upikoth/starter-go/internal/service/oauth"
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
	Oauth                    *oauth.Oauth
}

func New(
	log logger.Logger,
	cfg *config.Config,
	repo *repository.Repository,
) (*Service, error) {
	return &Service{
		Registrations: registrations.New(
			log,
			&cfg.Service.Registrations,
			repo,
		),
		Sessions: sessions.New(
			log,
			repo,
		),
		PasswordRecoveryRequests: passwordrecoveryrequests.New(
			log,
			&cfg.Service.PasswordRecoveryRequests,
			repo,
		),
		Users: users.New(
			log,
			repo,
		),
		Oauth: oauth.New(
			log,
			cfg.Service.Oauth,
			repo,
		),
	}, nil
}
