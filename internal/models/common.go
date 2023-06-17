package models

type HomeworkType int64

const (
	// since iota starts with 0, the first value
	// defined here will be the default
	_ HomeworkType = iota
	First
	Second
)
