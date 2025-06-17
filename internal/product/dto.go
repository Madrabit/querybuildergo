package product

import "database/sql"

type Response struct {
	ProductName []string `json:"productName"`
}

type Entity struct {
	Name sql.NullString
}

func ToResponse(entities []Entity) Response {
	names := make([]string, 0, len(entities))
	for _, entity := range entities {
		if entity.Name.Valid {
			names = append(names, entity.Name.String)
		}
	}
	return Response{ProductName: names}
}
