package employee

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"querybuilder/internal/common"
)

type Service struct {
	repo Repo
	gen  FileGenerator
}

type Repo interface {
	FindByProducts(tx *sqlx.Tx, products []string) ([]Entity, error)
	BeginTransaction() (tx *sqlx.Tx, err error)
	SetAnsiNullsOffTx(tx *sqlx.Tx) error
}

type FileGenerator interface {
	CreateExl(empls []Entity) ([]byte, error)
}

func NewService(repo Repo, gen FileGenerator) *Service {
	return &Service{repo, gen}
}

func (s *Service) FindByProducts(products []string) (file []byte, err error) {
	tx, err := s.repo.BeginTransaction()
	if err != nil {
		return nil, fmt.Errorf("employee service: find by products: error starting transaction")
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
		return nil, fmt.Errorf("employee service: find by products:  error set ansi null off: %w", err)
	}
	prod, err := s.repo.FindByProducts(tx, products)
	if err != nil {
		return nil, &common.NotFoundError{Message: "employee service: find by products: employees not found"}
	}
	exl, err := s.gen.CreateExl(prod)
	if err != nil {
		return nil, fmt.Errorf("employee service: find by products: error creating exl: %w", err)
	}
	return exl, nil
}
