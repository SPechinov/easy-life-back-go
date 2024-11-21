package group_users

import (
	"database/sql"
	"time"
)

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
