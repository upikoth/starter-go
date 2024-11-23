package oauth

import (
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/services/sessions"
	"github.com/upikoth/starter-go/internal/services/users"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"
)

type oauthServices struct {
	users    *users.Users
	sessions *sessions.Sessions
}

type Oauth struct {
	logger   logger.Logger
	vkConfig *oauth2.Config
	services *oauthServices
}

func New(
	log logger.Logger,
	cfg config.Oauth,
	usersSrv *users.Users,
	sessionsSrv *sessions.Sessions,
) *Oauth {
	return &Oauth{
		logger: log,
		vkConfig: &oauth2.Config{
			ClientID:     cfg.VkClientID,
			ClientSecret: cfg.VkClientSecret,
			Endpoint:     vk.Endpoint,
			RedirectURL:  cfg.VkRedirectURL,
			Scopes:       []string{"email"},
		},
		services: &oauthServices{
			users:    usersSrv,
			sessions: sessionsSrv,
		},
	}
}
