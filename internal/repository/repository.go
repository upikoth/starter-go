package repository

import (
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	ycpstarter "github.com/upikoth/starter-go/internal/repository/ycp-starter"
	ydbstarter "github.com/upikoth/starter-go/internal/repository/ydb-starter"
)

type Repository struct {
	YcpStarter *ycpstarter.YcpStarter
	YdbStarter *ydbstarter.YdbStarter
}

func New(
	logger logger.Logger,
	config *config.Repository,
) (*Repository, error) {
	ycpStarter, err := ycpstarter.New(logger, &config.YcpStarter)

	if err != nil {
		return nil, err
	}

	ydbStarter, err := ydbstarter.New(logger, &config.YdbStarter)

	if err != nil {
		return nil, err
	}

	return &Repository{
		YcpStarter: ycpStarter,
		YdbStarter: ydbStarter,
	}, nil
}

func (r *Repository) Connect() error {
	err := r.YdbStarter.Connect()

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Disconnect() error {
	return r.YdbStarter.Disconnect()
}
