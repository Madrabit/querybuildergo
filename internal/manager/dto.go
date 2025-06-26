package manager

import (
	"database/sql"
)

type Entity struct {
	FullName      sql.NullString `db:"fullname"`
	Dep           sql.NullString `db:"dep"`
	Bank          sql.NullString `db:"bank"`
	Comment       sql.NullString `db:"comment"`
	CommentSecond sql.NullString `db:"commentSecond"`
}

type DailyReportDTOReq struct {
	Manager   string `json:"manager,omitempty"`
	StartDate string `json:"startDate,omitempty"`
	EndDate   string `json:"endDate,omitempty"`
}

type Response struct {
	ManagerReport []CallReportMapped `json:"managerReport"`
}

type CallReportMapped struct {
	FullName      string
	Dep           string
	Bank          string
	Comment       string
	CommentSecond string
}

func (r *Entity) ToCallReport() CallReportMapped {
	return CallReportMapped{
		FullName:      StringOrEmpty(r.FullName),
		Dep:           StringOrEmpty(r.Dep),
		Bank:          StringOrEmpty(r.Bank),
		Comment:       StringOrEmpty(r.Comment),
		CommentSecond: StringOrEmpty(r.CommentSecond),
	}
}

func ToResponse(entities []Entity) Response {
	reports := make([]CallReportMapped, 0, len(entities))
	for _, entity := range entities {
		reports = append(reports, entity.ToCallReport())
	}
	return Response{ManagerReport: reports}
}

func StringOrEmpty(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}
