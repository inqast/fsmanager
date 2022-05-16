package db

import (
	"context"

	"github.com/inqast/fsmanager/internal/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

func New(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	return pgxpool.Connect(ctx, cfg.Database.GetConnString())
}
