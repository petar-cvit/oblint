package models

type Stats struct {
	GenerationScore []int   `json:"generationScore"`
	Min             float64 `json:"min"`
	Max             float64 `json:"max"`
	Avg             float64 `json:"avg"`
	Std             float64 `json:"std"`

	Finished   int `json:"finished"`
	InProgress int `json:"inProgress"`
	NotStarted int `json:"notStarted"`
}
