package repository

import (
	"context"

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

func (r *Repository) Connect(ctx context.Context) error {
	return r.YDB.Connect(ctx)
}

func (r *Repository) Disconnect(ctx context.Context) error {
	return r.YDB.Disconnect(ctx)
}
