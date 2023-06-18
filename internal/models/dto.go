package models

type SubmitHomeworkDTO struct {
	ID              string   `json:"id"`
	DataAnswer      []string `json:"dataAnswer"`
	StatementAnswer string   `json:"statementAnswer"`
}
