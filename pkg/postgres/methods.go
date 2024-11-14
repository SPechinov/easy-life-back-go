package postgres

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (p *Postgres) Exec(sql string, arguments ...any) (pgconn.CommandTag, error) {
	return p.pool.Exec(p.ctx, sql, arguments...)
}

func (p *Postgres) Query(sql string, arguments ...any) (pgx.Rows, error) {
	return p.pool.Query(p.ctx, sql, arguments...)
}

func (p *Postgres) QueryRow(sql string, arguments ...any) pgx.Row {
	return p.pool.QueryRow(p.ctx, sql, arguments...)
}
