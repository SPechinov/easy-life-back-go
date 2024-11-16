package composites

import (
	"context"
	"go-clean/pkg/postgres"
)

func NewPostgres(ctx context.Context, options *postgres.Options) (postgres.Client, error) {
	return postgres.New(ctx, options)
}
