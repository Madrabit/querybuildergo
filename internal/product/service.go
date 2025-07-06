package product

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"querybuilder/internal/common"
)

type Service struct {
	repo      Repo
	validator Validator
}

type Repo interface {
	GetAllProductsTx(tx *sqlx.Tx) ([]Entity, error)
	BeginTransaction() (tx *sqlx.Tx, err error)
	SetAnsiNullsOffTx(tx *sqlx.Tx) error
}

func NewService(repo Repo, validator Validator) *Service {
	return &Service{repo: repo, validator: validator}
}

type Validator interface {
	Validate(request any) error
}

func (s *Service) GetAllProducts() (products Response, err error) {
	tx, err := s.repo.BeginTransaction()
	if err != nil {
		return Response{}, fmt.Errorf("product service: get all products: error starting transaction")
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			err = fmt.Errorf("product service: get all products: panic getting products: %v", p)
			return
		}
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("rollback failed: original error: %w", err)
			}
			return
		}
		if commitErr := tx.Commit(); commitErr != nil {
			err = fmt.Errorf("product service: get all products: committing transaction failed: %w", commitErr)
		}
	}()
	err = s.repo.SetAnsiNullsOffTx(tx)
	if err != nil {
		return Response{}, fmt.Errorf("product service: get all products: error set ansi null off: %w", err)
	}
	prod, err := s.repo.GetAllProductsTx(tx)
	if err != nil {
		return Response{}, &common.NotFoundError{Message: "product service: get all products: products not found"}
	}
	products = ToResponse(prod)
	return products, nil
}
