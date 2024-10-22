package db

import (
	"easy-life-back-go/internal/db/users"
	"easy-life-back-go/pkg/postgres"
)

type DB struct {
	Users *users.Users
}

func New(postgres postgres.Postgres) *DB {
	return &DB{
		Users: users.New(postgres),
	}
}
