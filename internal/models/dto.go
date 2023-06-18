package models

type SubmitHomeworkDTO struct {
	ID              string `json:"id"`
	DataAnswer      []int  `json:"dataAnswer"`
	StatementAnswer []int  `json:"statementAnswer"`
}
