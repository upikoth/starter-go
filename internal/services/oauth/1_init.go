package oauth

import (
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repositories/http/oauth"
	"github.com/upikoth/starter-go/internal/services/sessions"
	"github.com/upikoth/starter-go/internal/services/users"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/mailru"
	"golang.org/x/oauth2/vk"
)

type oauthRepositories struct {
	oauth *oauth.Oauth
}

type oauthServices struct {
	users    *users.Users
	sessions *sessions.Sessions
}

type Oauth struct {
	logger       logger.Logger
	vkConfig     *oauth2.Config
	mailConfig   *oauth2.Config
	repositories *oauthRepositories
	services     *oauthServices
}

func New(
	log logger.Logger,
	cfg config.Oauth,
	oauthRepo *oauth.Oauth,
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
		mailConfig: &oauth2.Config{
			ClientID:     cfg.MailClientID,
			ClientSecret: cfg.MailClientSecret,
			Endpoint:     mailru.Endpoint,
			RedirectURL:  cfg.MailRedirectURL,
			Scopes:       []string{"userinfo"},
		},
		repositories: &oauthRepositories{
			oauth: oauthRepo,
		},
		services: &oauthServices{
			users:    usersSrv,
			sessions: sessionsSrv,
		},
	}
}
