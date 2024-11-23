package repositories

import (
	"context"

	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repositories/ycp"
	"github.com/upikoth/starter-go/internal/repositories/ydb"
)

type Repository struct {
	YCP *ycp.Ycp
	YDB *ydb.YDB
}

func New(
	logger logger.Logger,
	config *config.Repositories,
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
	return r.YDB.Migrate(ctx)
}

func (r *Repository) Disconnect(ctx context.Context) error {
	return r.YDB.Disconnect(ctx)
}
