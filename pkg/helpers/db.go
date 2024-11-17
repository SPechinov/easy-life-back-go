package helpers

import (
	"database/sql"
	"time"
)

func GetValueFromSQLNullString(data sql.NullString) string {
	if data.Valid {
		return data.String
	}

	return ""
}

func GetPtrValueFromSQLNullString(data sql.NullString) *string {
	if data.Valid {
		return StrToPtr(data.String)
	}

	return nil
}

func GetPtrValueFromSQLNullTime(data sql.NullTime, format string) *string {
	if data.Valid {
		StrToPtr(data.Time.Format(time.RFC3339))
	}

	return nil
}
