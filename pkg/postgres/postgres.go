package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	ctx              context.Context
	pool             *pgxpool.Pool
	ConnectionString string
}

type Options struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  bool
}

func New(ctx context.Context, options *Options) (*Postgres, error) {
	connectionString := getConnectionString(options)
	pool, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		return nil, fmt.Errorf("could not create postgres pool %s: %v", connectionString, err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database %s: %v", connectionString, err)
	}

	return &Postgres{ctx: ctx, pool: pool, ConnectionString: connectionString}, nil
}

func getConnectionString(options *Options) string {
	connectionString := "postgres://" + options.User + ":" + options.Password + "@" + options.Host + ":" + options.Port + "/" + options.DBName
	if !options.SSLMode {
		connectionString += "?sslmode=disable"
	}
	return connectionString
}

func (p *Postgres) Close() {
	p.pool.Close()
}
