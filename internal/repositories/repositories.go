package repositories

import (
	"context"

	"github.com/upikoth/starter-go/internal/config"
	"github.com/upikoth/starter-go/internal/pkg/logger"
	"github.com/upikoth/starter-go/internal/repositories/ydb"
)

type Repository struct {
	YDB *ydb.YDB
}

func New(
	log logger.Logger,
	cfg *config.Repositories,
) (*Repository, error) {
	ydbInstance, err := ydb.New(log, &cfg.Ydb)
	if err != nil {
		return nil, err
	}

	return &Repository{
		YDB: ydbInstance,
	}, nil
}

func (r *Repository) Connect(ctx context.Context) error {
	return r.YDB.Migrate(ctx)
}

func (r *Repository) Disconnect(ctx context.Context) error {
	return r.YDB.Disconnect(ctx)
}
