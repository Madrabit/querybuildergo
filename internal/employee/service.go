package employee

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"querybuilder/internal/common"
)

type Service struct {
	repo Repo
}

type Repo interface {
	FindByProducts(tx *sqlx.Tx, products []string) ([]Entity, error)
	BeginTransaction() (tx *sqlx.Tx, err error)
	SetAnsiNullsOffTx(tx *sqlx.Tx) error
}

func NewService(repo Repo) *Service {
	return &Service{repo}
}

func (s *Service) FindByProducts() (products Response, err error) {
	tx, err := s.repo.BeginTransaction()
	if err != nil {
		return Response{}, fmt.Errorf("employee service: find by products: error starting transaction")
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			err = fmt.Errorf("employee service: find by products: panic getting products: %v", p)
			return
		}
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("rollback failed: original error: %w", err)
			}
			return
		}
		if commitErr := tx.Commit(); commitErr != nil {
			err = fmt.Errorf("employee service: find by products:  committing transaction failed: %w", commitErr)
		}
	}()
	err = s.repo.SetAnsiNullsOffTx(tx)
	if err != nil {
		return Response{}, fmt.Errorf("employee service: find by products:  error set ansi null off: %w", err)
	}
	prod, err := s.repo.GetAllProductsTx(tx)
	if err != nil {
		return Response{}, &common.NotFoundError{Massage: "employee service: find by products: employees not found"}
	}
	products = ToResponse(prod)
	return products, nil
}
