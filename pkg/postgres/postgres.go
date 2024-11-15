package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"go-clean/pkg/helpers"
	"time"
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

	pool, err := connect(ctx, connectionString)
	if err != nil {
		return nil, err
	}

	return &Postgres{ctx: ctx, pool: pool, ConnectionString: connectionString}, nil
}

func connect(ctx context.Context, connectionString string) (*pgxpool.Pool, error) {
	var pool *pgxpool.Pool

	fmt.Println("Postgres connecting...")
	err := helpers.Repeatable(func() error {
		fmt.Println("Postgres try to connect")
		pool, poolErr := pgxpool.New(ctx, connectionString)
		if poolErr != nil {
			return poolErr
		}

		pingErr := pool.Ping(ctx)
		if pingErr != nil {
			return pingErr
		}

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
