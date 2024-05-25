package repository

import (
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	ycpstarter "github.com/upikoth/starter-go/internal/repository/ycp-starter"
)

type Repository struct {
	YcpStarter *ycpstarter.YcpStarter
	logger     logger.Logger
}

func New(
	logger logger.Logger,
	config *config.Repository,
) (*Repository, error) {
	ycpStarter, err := ycpstarter.New(logger, &config.YcpStarter)

	if err != nil {
		return nil, err
	}

	return &Repository{
		YcpStarter: ycpStarter,
		logger:     logger,
	}, nil
}
