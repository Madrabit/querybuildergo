package product

import "querybuilder/types"

type ProductName struct {
	ProductName []string `json:"productName"`
}

func ToProductDto(p *types.Product) string {
	return StringOrEmpty(p.Name)
}
