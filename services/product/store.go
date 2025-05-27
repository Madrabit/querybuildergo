package product

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"querybuilder/types"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetAllProducts(ctx context.Context) ([]*types.Product, error) {
	conn, err := s.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	_, err = conn.ExecContext(ctx, "set ansi_nulls off\n")
	if err != nil {
		return nil, err
	}
	rows, err := conn.QueryContext(ctx, "SELECT DISTINCT P794 as productName FROM dbo.Attr143 ORDER BY P794 ASC")
	if err != nil {
		return nil, fmt.Errorf("select Products error %v", err)
	}
	defer rows.Close()
	products := make([]*types.Product, 0)
	for rows.Next() {
		pr, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, fmt.Errorf("Error scanning row %v", err)
		}
		products = append(products, pr)
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error scanning rows %v", err)
	}
	return products, nil

}

func scanRowIntoProduct(rows *sql.Rows) (*types.Product, error) {
	prod := new(types.Product)
	err := rows.Scan(
		&prod.Name,
	)
	if err != nil {
		return nil, err
	}
	return prod, nil
}
