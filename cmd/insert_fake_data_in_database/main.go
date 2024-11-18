package main

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"go-clean/config"
	"go-clean/internal/composites"
	"go-clean/pkg/postgres"
	"strconv"
	"time"
)

var ctx = context.Background()

func main() {
	// Config
	cfg, err := config.InitConfig("./env.yaml")
	if err != nil {
		panic("Config not initialized")
	}

	// Postgres
	db, err := composites.NewPostgres(ctx, &postgres.Options{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		User:     cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		DBName:   cfg.Postgres.DBName,
		SSLMode:  cfg.Postgres.SSLMode,
	})

	users(ctx, db)
}

func users(ctx context.Context, postgres postgres.Client) {
	var count = 90000000
	start := time.Now()
	cursor := 0
	for cursor < count {
		email := gofakeit.Email()
		firstName := gofakeit.FirstName()
		lastName := gofakeit.LastName()
		password := "$2a$08$1HvQReL31t3g4xIraZwGXOQtOvvJnhEdwi.lH3FJbRgmrjWzALblG"

		query := `
			INSERT INTO public.users (email, first_name, last_name, password)
			VALUES ($1, $2, $3, $4)
		`

		_, err := postgres.Exec(ctx, query, strconv.Itoa(cursor)+email, firstName, lastName, password)
		if err != nil {
			panic(err)
		}

		cursor++

		if cursor%10000 == 0 {
			fmt.Println("Left: ", count-cursor)
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Время выполнения функции: %s\n", elapsed)
}
