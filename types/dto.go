package types

import (
	"database/sql"
)

type EmployeeDTOResp struct {
	//attr143.P794
	ProductName sql.NullString
	//P18
	ShortBankName sql.NullString
	// P2372
	FullBankName sql.NullString
	//P1972
	LastName sql.NullString
	//P1973
	Name sql.NullString
	//P1974
	Patronymic sql.NullString
	//P948
	JobTitle sql.NullString
	//P971
	Email sql.NullString
	//attr166.P3224
	ContactDate sql.NullTime
	//attr166.p970,
	Phone sql.NullString
	//attr166.p2424,
	ExtensionPhone sql.NullString
	//attr166.p1975
	Mobile sql.NullString
}
type CallReport struct {
	FullName      sql.NullString
	Dep           sql.NullString
	Bank          sql.NullString
	Comment       sql.NullString
	CommentSecond sql.NullString
}
