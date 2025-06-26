package employee

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

func (r *Repository) FindByProducts(tx *sqlx.Tx, products []string) (entity []Entity, err error) {
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
	if err = tx.Select(&entity, query, args...); err != nil {
		return nil, err
	}
	return entity, nil
}
