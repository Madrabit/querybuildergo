package manager

import "database/sql"

type CallReport struct {
	FullName      sql.NullString
	Dep           sql.NullString
	Bank          sql.NullString
	Comment       sql.NullString
	CommentSecond sql.NullString
}

type DailyReportDTOReq struct {
	Manager   string `json:"manager,omitempty"`
	StartDate string `json:"startDate,omitempty"`
	EndDate   string `json:"endDate,omitempty"`
}

type ManagerReportResp struct {
	ManagerReport []CallReportMapped `json:"managerReport"`
}

type CallReportMapped struct {
	FullName      string
	Dep           string
	Bank          string
	Comment       string
	CommentSecond string
}

func (r *CallReport) ToCallReport() CallReportMapped {
	return CallReportMapped{
		FullName:      StringOrEmpty(r.FullName),
		Dep:           StringOrEmpty(r.Dep),
		Bank:          StringOrEmpty(r.Bank),
		Comment:       StringOrEmpty(r.Comment),
		CommentSecond: StringOrEmpty(r.CommentSecond),
	}
}
