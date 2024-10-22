package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres interface {
	Exec(query string, args ...interface{}) (pgconn.CommandTag, error)
	Close()
}

type postgres struct {
	Pool *pgxpool.Pool
}

var ctx = context.Background()

func New(connStr string) (Postgres, error) {
	connection, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, errors.New("failed to connect to postgres: " + err.Error())
	}

	return &postgres{Pool: connection}, nil
}

func (p *postgres) Exec(query string, args ...interface{}) (pgconn.CommandTag, error) {
	res, err := p.Pool.Exec(ctx, query, args...)

	if err != nil {
		return res, errors.New("failed to execute query: " + err.Error())
	}

	return res, nil
}

func (p *postgres) Close() {
	p.Pool.Close()
}
