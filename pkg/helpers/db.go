package helpers

import (
	"database/sql"
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
		return StrToPtr(data.Time.Format(format))
	}

	return nil
}
