package clients

import (
	httpclient "github.com/upikoth/starter-go/internal/clients/http-client"
	ycpclient "github.com/upikoth/starter-go/internal/clients/ycp-client"
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"go.opentelemetry.io/otel/trace"
)

type Clients struct {
	YCP  *ycpclient.Ycp
	HTTP *httpclient.HTTP
}

func New(
	log logger.Logger,
	cfg *config.Clients,
	tp trace.TracerProvider,
) (*Clients, error) {
	ycpInstance, err := ycpclient.New(log, &cfg.Ycp)
	if err != nil {
		return nil, err
	}

	httpInstance, err := httpclient.New(log, &cfg.HTTP, tp)
	if err != nil {
		return nil, err
	}

	return &Clients{
		YCP:  ycpInstance,
		HTTP: httpInstance,
	}, nil
}
