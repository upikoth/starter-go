package services

import (
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repositories"
	"github.com/upikoth/starter-go/internal/services/emails"
	"github.com/upikoth/starter-go/internal/services/oauth"
	passwordrecoveryrequests "github.com/upikoth/starter-go/internal/services/password-recovery-requests"
	"github.com/upikoth/starter-go/internal/services/registrations"
	"github.com/upikoth/starter-go/internal/services/sessions"
	"github.com/upikoth/starter-go/internal/services/users"
)

type Services struct {
	Registrations            *registrations.Registrations
	Sessions                 *sessions.Sessions
	PasswordRecoveryRequests *passwordrecoveryrequests.PasswordRecoveryRequests
	Users                    *users.Users
	Oauth                    *oauth.Oauth
	Emails                   *emails.Emails
}

func New(
	log logger.Logger,
	cfg *config.Config,
	repo *repositories.Repository,
) (*Services, error) {
	srvs := &Services{}

	srvs.Emails = emails.New(
		log,
		&cfg.Services.Emails,
		repo.YCP,
	)

	srvs.Users = users.New(
		log,
		repo.YDB.DB,
		repo.YDB.Users,
	)

	srvs.Sessions = sessions.New(
		log,
		repo.YDB.DB,
		repo.YDB.Sessions,
		srvs.Users,
	)

	srvs.Registrations = registrations.New(
		log,
		repo.YDB.DB,
		repo.YDB.Registrations,
		srvs.Users,
		srvs.Sessions,
		srvs.Emails,
	)

	srvs.PasswordRecoveryRequests = passwordrecoveryrequests.New(
		log,
		repo.YDB.DB,
		repo.YDB.PasswordRecoveryRequests,
		srvs.Users,
		srvs.Sessions,
		srvs.Emails,
	)

	srvs.Oauth = oauth.New(
		log,
		cfg.Services.Oauth,
		srvs.Users,
		srvs.Sessions,
	)

	return srvs, nil
}
