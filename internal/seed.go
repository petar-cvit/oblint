package internal

import "example.com/oblint/internal/models"

const (
	firstQuestion  = "Write down the statement for the boolean data"
	secondQuestion = "Fill in the empty Res fields for the boolean table"
)

func seedHistory(storage Storage) {
	err := storage.SaveToHistory(models.HistoryHomework{
		ID:             "123",
		Name:           "Prva zadaca",
		SubmissionDate: "12.5.2023.",
		DueDate:        "11.5.2023.",
		Points:         5,
		MaxPoints:      10,
		Type:           models.Second,
		Question:       secondQuestion,
		Statement:      "A and B",
		Data:           []string{"0", "0", "0", "0", "1", "0", "1", "0", "1", "1", "1", "1"},
		CorrectData:    []string{"0", "0", "0", "0", "1", "0", "1", "0", "0", "1", "1", "1"},
		Answer:         "",
		CorrectAnswer:  "",
	})
	if err != nil {
		panic(err)
	}

	err = storage.SaveToHistory(models.HistoryHomework{
		ID:             "124",
		Name:           "Druga zadaca",
		SubmissionDate: "15.5.2023.",
		DueDate:        "11.5.2023.",
		Points:         9,
		MaxPoints:      12,
		Type:           models.Second,
		Question:       secondQuestion,
		Statement:      "A and not B",
		Data:           []string{"0", "0", "0", "0", "1", "0", "1", "0", "0", "1", "1", "1"},
		CorrectData:    []string{"0", "0", "0", "0", "1", "0", "1", "0", "1", "1", "1", "0"},
		Answer:         "",
		CorrectAnswer:  "",
	})
	if err != nil {
		panic(err)
	}

	err = storage.SaveToHistory(models.HistoryHomework{
		ID:             "125",
		Name:           "Treca zadaca",
		SubmissionDate: "15.5.2023.",
		DueDate:        "14.5.2023.",
		Points:         1,
		MaxPoints:      5,
		Type:           models.First,
		Question:       firstQuestion,
		Statement:      "",
		CorrectData:    []string{"0", "0", "0", "0", "1", "1", "1", "0", "1", "1", "1", "1"},
		Data:           []string{},
		Answer:         "not A and B",
		CorrectAnswer:  "A or B",
	})
	if err != nil {
		panic(err)
	}
}

func seedOngoing(storage Storage) {
	err := storage.SaveToHomeworks(models.Homework{
		ID:                 "321",
		Name:               "Prva zadaca za napraviti",
		LastSubmissionDate: "",
		DueDate:            "15.5.2023.",
		MaxPoints:          10,
		Type:               models.First,
		Question:           firstQuestion,
		Statement:          "",
		CorrectData:        []string{"0", "0", "0", "0", "1", "1", "1", "0", "1", "1", "1", "1"},
		Data:               []string{},
		Answer:             "",
		CorrectAnswer:      "A or not B",
		Started:            false,
	})
	if err != nil {
		panic(err)
	}

	err = storage.SaveToHomeworks(models.Homework{
		ID:                 "322",
		Name:               "Prva zadaca za napraviti",
		LastSubmissionDate: "15.5.2023.",
		DueDate:            "15.5.2023.",
		MaxPoints:          8,
		Type:               models.Second,
		Question:           secondQuestion,
		Statement:          "A and not B",
		Data:               []string{"0", "0", "0", "0", "1", "1", "1", "0", "", "1", "1", "1"},
		CorrectData:        []string{"0", "0", "0", "0", "1", "1", "1", "0", "1", "1", "1", "1"},
		Answer:             "",
		CorrectAnswer:      "",
		Started:            true,
	})
	if err != nil {
		panic(err)
	}

	err = storage.SaveToHomeworks(models.Homework{
		ID:                 "323",
		Name:               "Treca zadaca za napraviti",
		LastSubmissionDate: "9.5.2023.",
		DueDate:            "12.5.2023.",
		MaxPoints:          16,
		Type:               models.First,
		Question:           firstQuestion,
		CorrectData:        []string{"0", "0", "0", "0", "1", "1", "1", "0", "1", "1", "1", "1"},
		Data:               []string{},
		Answer:             "",
		CorrectAnswer:      "A and B",
		Started:            false,
	})
	if err != nil {
		panic(err)
	}
}
