package users

import (
	"database/sql"
	"time"
)

type dataUser struct {
	ID        string
	Email     sql.NullString
	Phone     sql.NullString
	Password  sql.RawBytes
	FirstName string
	LastName  sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
