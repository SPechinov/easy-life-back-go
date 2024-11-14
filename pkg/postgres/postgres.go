package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
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
	fmt.Println("Postgres connecting...")
	pool, connectionErr := pgxpool.New(ctx, connectionString)

	if connectionErr == nil {
		connectionErr = pool.Ping(ctx)

		if connectionErr == nil {
			fmt.Println("Postgres connected")
			return pool, nil
		}
	}

	tryCount := 1
	for tryCount < 10 {
		time.Sleep(2 * time.Second)

		fmt.Printf("Postgres try to connect: %d time\n", tryCount+1)
		tryCount++

		pool, connectionErr = pgxpool.New(ctx, connectionString)
		if connectionErr == nil {
			connectionErr = pool.Ping(ctx)

			if connectionErr == nil {
				fmt.Println("Postgres connected")
				return pool, nil
			}
		}
	}

	fmt.Printf("Postgres not connected: %s\n", connectionErr)
	return nil, connectionErr
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
