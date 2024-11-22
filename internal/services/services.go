package services

import (
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repositories"
	"github.com/upikoth/starter-go/internal/services/oauth"
	passwordrecoveryrequests "github.com/upikoth/starter-go/internal/services/password-recovery-requests"
	"github.com/upikoth/starter-go/internal/services/registrations"
	"github.com/upikoth/starter-go/internal/services/sessions"
	"github.com/upikoth/starter-go/internal/services/users"
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
	repo *repositories.Repository,
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
