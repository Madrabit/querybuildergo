package product

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) BeginTransaction() (tx *sqlx.Tx, err error) {
	return r.db.Beginx()
}

func (r *Repository) SetAnsiNullsOffTx(tx *sqlx.Tx) error {
	_, err := tx.Exec("set ansi_nulls off\n")
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetAllProductsTx(tx *sqlx.Tx) ([]Entity, error) {
	var entities []Entity
	err := tx.Select(&entities, "SELECT DISTINCT P794 as name FROM dbo.Attr143 ORDER BY P794 ASC")
	if err != nil {
		return nil, err
	}
	return entities, nil
}
