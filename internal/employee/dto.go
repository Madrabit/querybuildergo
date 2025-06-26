package employee

import "database/sql"

type EmployeeClean struct {
	ProductName    string
	ShortBankName  string
	FullBankName   string
	LastName       string
	Name           string
	Patronymic     string
	JobTitle       string
	Email          string
	ContactDate    string
	Phone          string
	ExtensionPhone string
	Mobile         string
}

type ProductsReq struct {
	Products []string `json:"products"`
}

type Entity struct {
	ProductName    sql.NullString `db:"productName"`
	ShortBankName  sql.NullString `db:"shortBankName"`
	FullBankName   sql.NullString `db:"fullBankName"`
	LastName       sql.NullString `db:"lastName"`
	Name           sql.NullString `db:"name"`
	Patronymic     sql.NullString `db:"patronymic"`
	JobTitle       sql.NullString `db:"jobTitle"`
	Email          sql.NullString `db:"email"`
	ContactDate    sql.NullTime   `db:"contactDate"`
	Phone          sql.NullString `db:"phone"`
	ExtensionPhone sql.NullString `db:"extensionPhone"`
	Mobile         sql.NullString `db:"mobile"`
}

func (e Entity) ToCallReport() EmployeeClean {
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

func StringOrEmpty(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}

func TimeOrEmpty(t sql.NullTime, layout string) string {
	if t.Valid {
		return t.Time.Format(layout)
	}
	return ""
}
