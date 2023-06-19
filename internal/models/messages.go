package models

type Message struct {
	Name          string `json:"name"`
	Timestamp     string `json:"timestamp"`
	MessageBody   string `json:"messageBody"`
	IsCurrentUser bool   `json:"isCurrentUser"`
}
