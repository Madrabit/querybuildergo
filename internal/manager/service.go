package manager

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"querybuilder/internal/common"
)

type Service struct {
	repo Repo
}

type Repo interface {
	GetDailyReport(tx *sqlx.Tx, manager, startDate, endDate string) ([]Entity, error)
	BeginTransaction() (tx *sqlx.Tx, err error)
	SetAnsiNullsOffTx(tx *sqlx.Tx) error
}

func NewService(repo Repo) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetDailyReport(manager, startDate, endDate string) (Response, error) {
	tx, err := s.repo.BeginTransaction()
	if err != nil {
		return Response{}, fmt.Errorf("manager service: get daily report: error starting transaction")
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			err = fmt.Errorf("manager service: get daily report: panic while getting report: %v", p)
			return
		}
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("rollback failed: original error: %w", err)
			}
			return
		}
		if commitErr := tx.Commit(); commitErr != nil {
			err = fmt.Errorf("manager service: get daily report: committing transaction failed: %w", commitErr)
		}
	}()
	err = s.repo.SetAnsiNullsOffTx(tx)
	if err != nil {
		return Response{}, fmt.Errorf("manager service: get daily report: error set ansi null off: %w", err)
	}
	report, err := s.repo.GetDailyReport(tx, manager, startDate, endDate)
	if err != nil {
		return Response{}, fmt.Errorf("manager service: get daily report: error: %w", err)
	}
	resp := ToResponse(report)
	if len(report) == 0 {
		return resp, &common.NotFoundError{Message: "manager service: get daily report: report not found"}
	}
	return resp, nil
}
