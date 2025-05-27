package employee

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{db: db}
}

func (s *Store) findByProducts(ctx context.Context, products []string) ([]EmployeeDTOResp, error) {
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
	if len(products) == 0 {
		return nil, nil
	}
	//TODO если привязать к бд, бд не поломается?
	query, args, err := sqlx.In(` 
		SELECT 
	P794 as productName,
		P18 as shortBankName,
		P2372 as fullBankName,
		P1972 as lastName,
		P1973 as name,
		P1974 as patronymic,
		P948 as jobTitle,
		P971 as email,
		P3224 as contactDate,
		p970 as phone,
		p2424 as extensionPhone,
		p1975 as mobile
	FROM dbo.Attr143 AS Product
	LEFT JOIN dbo.attr333 AS Shedule
	ON Product.ObjectID = Shedule.P2512
	LEFT JOIN dbo.attr355 AS clients
	ON Shedule.ObjectID = clients.P2749
	LEFT JOIN dbo.attr166 AS employ
	ON clients.P2696 = employ.ObjectID
	LEFT JOIN dbo.attr5 AS bank
	ON bank.ObjectID = employ.P944
	LEFT JOIN attr369
	ON attr369.ObjectID = employ.P2837
	WHERE
	P2838 IS NOT NULL AND
	P3177 IS NOT NULL AND
	P2385 IS NULL AND
	P2386 IS NULL AND
	(bank.p3284 = 0 or bank.p3284 is null) AND
	(bank.p2385 = 0 or bank.p2385 is null) AND
	(bank.p2386 = 0 or bank.p2386 is null) AND
	P794 IN (?)
	ORDER BY P794, P18, P1972 ASC`, products)
	if err != nil {
		return nil, err
	}
	query = tx.Rebind(query)
	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("select Products error %v", err)
	}
	defer rows.Close()
	empl := make([]EmployeeDTOResp, 0)
	for rows.Next() {
		e, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, fmt.Errorf("Error scanning row %v", err)
		}
		empl = append(empl, e)
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error scanning rows %v", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return empl, nil

}

func scanRowIntoProduct(rows *sql.Rows) (EmployeeDTOResp, error) {
	empl := EmployeeDTOResp{}
	err := rows.Scan(
		&empl.ProductName, &empl.ShortBankName, &empl.FullBankName, &empl.LastName, &empl.Name,
		&empl.Patronymic, &empl.JobTitle, &empl.Email, &empl.ContactDate, &empl.Phone,
		&empl.ExtensionPhone, &empl.Mobile,
	)
	if err != nil {
		return EmployeeDTOResp{}, err
	}
	return empl, nil
}
