package employee

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
