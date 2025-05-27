package product

import (
	"database/sql"
)

func StringOrEmpty(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}

func TimeOrEmpty(t sql.NullTime, layout string) string {
	if t.Valid {
		return t.Time.Format(layout)
	}
	return ""
}
