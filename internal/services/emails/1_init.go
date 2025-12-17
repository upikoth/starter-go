package emails

import (
	ycpclient "github.com/upikoth/starter-go/internal/clients/ycp-client"
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
)

type emailsClients struct {
	ycp *ycpclient.Ycp
}

type Emails struct {
	logger       logger.Logger
	config       *config.Emails
	repositories *emailsClients
}

func New(
	logger logger.Logger,
	cfg *config.Emails,
	ycp *ycpclient.Ycp,
) *Emails {
	return &Emails{
		logger: logger,
		config: cfg,
		repositories: &emailsClients{
			ycp: ycp,
		},
	}
}
