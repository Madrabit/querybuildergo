package manager

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) SetAnsiNullsOffTx(tx *sqlx.Tx) error {
	_, err := tx.Exec("set ansi_nulls off\n")
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) BeginTransaction() (tx *sqlx.Tx, err error) {
	return r.db.Beginx()
}

func (r *Repository) GetDailyReport(tx *sqlx.Tx, manager, startDate, endDate string) (report []Entity, err error) {
	query := `
		SELECT 
			a.P1972 + ' ' + a.P1973 + ' ' + a.P1974 as fullname, P2835 as dep, P2599 as comment,P3290 as commentSecond, P18 as bank
		FROM dbo.Attr166 a
			LEFT JOIN dbo.attr6 b ON  a.User_mod = b.ObjectID
			LEFT JOIN dbo.attr344 c ON a.ObjectID  = c.P2594 
			LEFT JOIN dbo.attr369 o ON a.P2837 = o.ObjectID
			LEFT JOIN dbo.attr5 bk ON a.P944 = bk.ObjectID
		WHERE ((c.Date_mod > ?  AND c.Date_mod < ?) 
			OR (c.Date_cr > ?  AND c.Date_cr  < ?))
			AND b.P37 Like ?
		ORDER BY c.Date_mod`
	args := []interface{}{startDate, endDate, startDate, endDate, "%" + manager + "%"}
	query = tx.Rebind(query)
	err = tx.Select(&report, query, args...)
	if err != nil {
		return nil, err
	}
	return report, nil
}
