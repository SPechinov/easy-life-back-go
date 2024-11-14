package composites

import (
	"context"
	"go-clean/pkg/postgres"
)

func NewPostgres(ctx context.Context, options *postgres.Options) (*postgres.Postgres, error) {
	return postgres.New(ctx, options)
}
