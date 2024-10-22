package users

import (
	"easy-life-back-go/pkg/postgres"
	"log/slog"
)

type Users struct {
	postgres postgres.Postgres
}

func New(postgres postgres.Postgres) *Users {
	u := &Users{postgres: postgres}
	u.init()
	return u
}

func (u *Users) init() {
	_, err := u.postgres.Exec(createTableSQL)
	if err != nil {
		slog.Error(err.Error())
	}
}
