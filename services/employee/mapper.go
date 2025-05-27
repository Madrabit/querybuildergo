package employee

import (
	"database/sql"
)

type EmployeeDTOResp struct {
	ProductName    sql.NullString
	ShortBankName  sql.NullString
	FullBankName   sql.NullString
	LastName       sql.NullString
	Name           sql.NullString
	Patronymic     sql.NullString
	JobTitle       sql.NullString
	Email          sql.NullString
	ContactDate    sql.NullTime
	Phone          sql.NullString
	ExtensionPhone sql.NullString
	Mobile         sql.NullString
}

func (e EmployeeDTOResp) ToCallReport() EmployeeClean {
	return EmployeeClean{
		ProductName:    StringOrEmpty(e.ProductName),
		ShortBankName:  StringOrEmpty(e.ShortBankName),
		FullBankName:   StringOrEmpty(e.FullBankName),
		LastName:       StringOrEmpty(e.LastName),
		Name:           StringOrEmpty(e.Name),
		Patronymic:     StringOrEmpty(e.Patronymic),
		JobTitle:       StringOrEmpty(e.JobTitle),
		Email:          StringOrEmpty(e.Email),
		ContactDate:    TimeOrEmpty(e.ContactDate, "2006-01-02"),
		Phone:          StringOrEmpty(e.Phone),
		ExtensionPhone: StringOrEmpty(e.ExtensionPhone),
		Mobile:         StringOrEmpty(e.Mobile),
	}
}
