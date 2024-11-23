package emails

import (
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repositories/ycp"
)

type emailsRepositories struct {
	ycp *ycp.Ycp
}

type Emails struct {
	logger       logger.Logger
	config       *config.Emails
	repositories *emailsRepositories
}

func New(
	logger logger.Logger,
	cfg *config.Emails,
	ycpRepo *ycp.Ycp,
) *Emails {
	return &Emails{
		logger: logger,
		config: cfg,
		repositories: &emailsRepositories{
			ycp: ycpRepo,
		},
	}
}
