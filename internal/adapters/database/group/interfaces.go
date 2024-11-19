package group

import (
	"database/sql"
	"time"
)

type dataGroup struct {
	id        string
	name      string
	isPayed   bool
	createdAt time.Time
	updatedAt time.Time
	deletedAt sql.NullTime
}

type dataUser struct {
	id         string
	email      sql.NullString
	phone      sql.NullString
	permission int
	firstName  string
	lastName   sql.NullString
	invitedAt  time.Time
	createdAt  time.Time
	updatedAt  time.Time
	deletedAt  sql.NullTime
}
