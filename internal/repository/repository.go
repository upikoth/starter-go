package repository

import (
	"github.com/upikoth/starter-go/internal/pkg/logger"
	ycpstarter "github.com/upikoth/starter-go/internal/repository/ycp-starter"
)

type Repository struct {
	YcpStarter *ycpstarter.YcpStarter
	logger     logger.Logger
}

func New(logger logger.Logger) (*Repository, error) {
	return &Repository{
		YcpStarter: ycpstarter.New(logger),
		logger:     logger,
	}, nil
}
