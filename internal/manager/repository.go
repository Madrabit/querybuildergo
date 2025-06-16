package manager

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (s *Repository) GetDailyReport(ctx context.Context, manager, startDate, endDate string) ([]CallReport, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = tx.Rollback()
	}()
	_, err = tx.ExecContext(ctx, "set ansi_nulls off\n")
	if err != nil {
		return nil, err
	}
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
	if err != nil {
		return nil, err
	}
	args := []interface{}{startDate, endDate, startDate, endDate, "%" + manager + "%"}
	query = tx.Rebind(query)
	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	callReport := make([]CallReport, 0)
	for rows.Next() {
		r, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, fmt.Errorf("error scanning row %v", err)
		}
		callReport = append(callReport, r)
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error scanning rows %v", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return callReport, nil
}

func scanRowIntoProduct(rows *sql.Rows) (CallReport, error) {
	report := CallReport{}
	err := rows.Scan(
		&report.FullName,
		&report.Dep,
		&report.Comment,
		&report.CommentSecond,
		&report.Bank,
	)
	if err != nil {
		return CallReport{}, err
	}
	return report, nil
}
