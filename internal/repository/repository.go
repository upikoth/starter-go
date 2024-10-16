package repository

import (
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository/ycp"
	"github.com/upikoth/starter-go/internal/repository/ydb"
)

type Repository struct {
	YCP *ycp.Ycp
	YDB *ydb.YDB
}

func New(
	logger logger.Logger,
	config *config.Repository,
) (*Repository, error) {
	ycpInstance, err := ycp.New(logger, &config.Ycp)

	if err != nil {
		return nil, err
	}

	ydbInstance, err := ydb.New(logger, &config.Ydb)

	if err != nil {
		return nil, err
	}

	return &Repository{
		YCP: ycpInstance,
		YDB: ydbInstance,
	}, nil
}

func (r *Repository) Connect() error {
	err := r.YDB.Connect()

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Disconnect() error {
	return r.YDB.Disconnect()
}
