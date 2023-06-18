package models

type HistoryHomework struct {
	ID             string       `json:"id"`
	Name           string       `json:"name"`
	SubmissionDate string       `json:"submissionDate"`
	DueDate        string       `json:"dueDate"`
	Points         int          `json:"points"`
	MaxPoints      int          `json:"maxPoints"`
	Type           HomeworkType `json:"type"`

	Question      string   `json:"question"`
	Statement     string   `json:"statement"`
	Data          []string `json:"data"`
	CorrectData   []string `json:"correctData"`
	Answer        string   `json:"answer"`
	CorrectAnswer string   `json:"correctAnswer"`
}

type Homework struct {
	ID                 string       `json:"id"`
	Name               string       `json:"name"`
	LastSubmissionDate string       `json:"lastSubmissionDate"`
	DueDate            string       `json:"dueDate"`
	MaxPoints          int          `json:"maxPoints"`
	Type               HomeworkType `json:"type"`
	Started            bool         `json:"started"`

	Question      string   `json:"question"`
	Statement     string   `json:"statement"`
	Data          []string `json:"data"`
	CorrectData   []string `json:"correctData"`
	Answer        string   `json:"answer"`
	CorrectAnswer string   `json:"correctAnswer"`
}
