package group

import (
	"database/sql"
	"time"
)

type dataGroup struct {
	ID        string
	Name      string
	IsPayed   bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
