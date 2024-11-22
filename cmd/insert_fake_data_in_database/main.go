package main

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"go-clean/config"
	"go-clean/internal/composites"
	"go-clean/pkg/helpers"
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
	_, err = composites.NewPostgres(ctx, &postgres.Options{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		User:     cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		DBName:   cfg.Postgres.DBName,
		SSLMode:  cfg.Postgres.SSLMode,
	})

	fmt.Println(helpers.GenerateRandomCode(6))
	fmt.Println(helpers.GenerateRandomCode(6))
	fmt.Println(helpers.GenerateRandomCode(6))
	fmt.Println(helpers.GenerateRandomCode(6))
	fmt.Println(helpers.GenerateRandomCode(6))
	fmt.Println(helpers.GenerateRandomCode(6))
	fmt.Println(helpers.GenerateRandomCode(6))
	fmt.Println(helpers.GenerateRandomCode(6))
	fmt.Println(helpers.GenerateRandomCode(6))
	fmt.Println(helpers.GenerateRandomCode(6))
	fmt.Println(helpers.GenerateRandomCode(6))
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

func addGroup(ctx context.Context, postgres postgres.Client) {
	tx, err := postgres.Begin(ctx)
	if err != nil {
		return
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	// Add group
	queryAddGroup := `
		INSERT INTO public.groups (name)
		VALUES ($1)
		RETURNING id
		`
	name := gofakeit.MovieName()

	var groupID string
	err = tx.QueryRow(ctx, queryAddGroup, name).Scan(&groupID)
	if err != nil {
		return
	}

	userId := "a026d602-9472-44d6-8742-ed611e8ef94e"
	queryUsersGroup := `
		INSERT INTO public.groups_users (group_id, user_id, permission)
		VALUES ($1, $2, $3)
	`
	_, err = tx.Exec(ctx, queryUsersGroup, groupID, userId, 777)
	if err != nil {
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		return
	}
}

func groups(ctx context.Context, postgres postgres.Client) {
	var count = 100000
	start := time.Now()
	cursor := 10
	for cursor < count {
		addGroup(ctx, postgres)
		cursor++

		if cursor%1000 == 0 {
			fmt.Println("Left: ", count-cursor)
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Время выполнения функции: %s\n", elapsed)
}

func addUserInGroup(ctx context.Context, postgres postgres.Client, groupID, userID string) {
	queryUsersGroup := `
			INSERT INTO public.groups_users (group_id, user_id, permission)
			VALUES ($1, $2, 0)
	`
	_, err := postgres.Exec(ctx, queryUsersGroup, groupID, userID)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func groupUsers(ctx context.Context, postgres postgres.Client) {
	start := time.Now()
	query := `
		SELECT id FROM public.users WHERE id != 'a026d602-9472-44d6-8742-ed611e8ef94e' ORDER BY created_at DESC LIMIT 100000 
	`

	ids := make([]string, 0)

	rows, err := postgres.Query(ctx, query)
	if err != nil {
		return
	}

	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			return
		}
		ids = append(ids, id)
	}

	for _, id := range ids {
		addUserInGroup(ctx, postgres, "3929910b-1069-4597-af8c-6c0c76d598fc", id)
	}
	elapsed := time.Since(start)
	fmt.Printf("Время выполнения функции: %s\n", elapsed)
}
