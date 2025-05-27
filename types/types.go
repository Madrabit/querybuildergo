package types

import "database/sql"

type Product struct {
	Name sql.NullString
}
