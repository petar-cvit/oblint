package models

type HistoryHomework struct {
	ID             string       `json:"id"`
	Name           string       `json:"name"`
	SubmissionDate string       `json:"submissionDate"`
	DueDate        string       `json:"dueDate"`
	Points         int          `json:"points"`
	MaxPoints      int          `json:"maxPoints"`
	Type           HomeworkType `json:"type"`
	Correct        string       `json:"correct"`
	Answered       string       `json:"answered"`
}

type Homework struct {
	ID                 string       `json:"id"`
	Name               string       `json:"name"`
	LastSubmissionDate string       `json:"lastSubmissionDate"`
	DueDate            string       `json:"dueDate"`
	MaxPoints          int          `json:"maxPoints"`
	Type               HomeworkType `json:"type"`
	Answered           string       `json:"answered"`
	Started            bool         `json:"started"`
}
