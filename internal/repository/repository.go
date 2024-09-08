package repository

import (
	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repository/ycp"
	"github.com/upikoth/starter-go/internal/repository/ydb"
)

type Repository struct {
	Ycp *ycp.Ycp
	Ydb *ydb.Ydb
}

func New(
	logger logger.Logger,
	config *config.Repository,
) (*Repository, error) {
	ycp, err := ycp.New(logger, &config.Ycp)

	if err != nil {
		return nil, err
	}

	ydb, err := ydb.New(logger, &config.Ydb)

	if err != nil {
		return nil, err
	}

	return &Repository{
		Ycp: ycp,
		Ydb: ydb,
	}, nil
}

func (r *Repository) Connect() error {
	err := r.Ydb.Connect()

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Disconnect() error {
	return r.Ydb.Disconnect()
}
