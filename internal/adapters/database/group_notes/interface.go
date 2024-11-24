package group_notes

import (
	"database/sql"
	"time"
)

type dataNote struct {
	id          string
	title       string
	description string
	groupID     string
	creatorID   string
	updaterID   string
	createdAt   time.Time
	updatedAt   time.Time
	deletedAt   sql.NullTime
}
