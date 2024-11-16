package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go-clean/pkg/helpers"
	"time"
)

type Client interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Close()
}

type Options struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  bool
}

func New(ctx context.Context, options *Options) (Client, error) {
	connectionString := getConnectionString(options)

	pool, err := connect(ctx, connectionString)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func connect(ctx context.Context, connectionString string) (*pgxpool.Pool, error) {
	var pool *pgxpool.Pool

	fmt.Println("Postgres connecting...")
	err := helpers.Repeatable(func() error {
		fmt.Println("Postgres try to connect")
		pl, poolErr := pgxpool.New(ctx, connectionString)
		if poolErr != nil {
			return poolErr
		}

		pingErr := pl.Ping(ctx)
		if pingErr != nil {
			return pingErr
		}

		pool = pl

		return nil
	}, 10, 2*time.Second)

	if err != nil {
		fmt.Printf("Postgres not connected: %s\n", err)
		return nil, err
	}

	fmt.Println("Postgres connected")
	return pool, nil
}

func getConnectionString(options *Options) string {
	connectionString := "postgres://" + options.User + ":" + options.Password + "@" + options.Host + ":" + options.Port + "/" + options.DBName
	if !options.SSLMode {
		connectionString += "?sslmode=disable"
	}
	return connectionString
}
