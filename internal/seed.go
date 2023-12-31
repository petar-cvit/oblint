package internal

import "example.com/oblint/internal/models"

const (
	FirstQuestion  = "Write down the statement for the boolean data"
	SecondQuestion = "Fill in the empty Res fields for the boolean table"
)

func seedHistory(storage Storage) {
	err := storage.SaveToHistory(models.HistoryHomework{
		ID:             "123",
		Name:           "Homework 1",
		SubmissionDate: "12.05.2023.",
		DueDate:        "11.05.2023.",
		Points:         5,
		MaxPoints:      10,
		Type:           models.Second,
		Question:       SecondQuestion,
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
		Name:           "Homework 2",
		SubmissionDate: "15.05.2023.",
		DueDate:        "11.05.2023.",
		Points:         9,
		MaxPoints:      10,
		Type:           models.Second,
		Question:       SecondQuestion,
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
		Name:           "Homework 3",
		SubmissionDate: "15.05.2023.",
		DueDate:        "14.05.2023.",
		Points:         1,
		MaxPoints:      10,
		Type:           models.First,
		Question:       FirstQuestion,
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
		Name:               "Homework 4",
		LastSubmissionDate: "",
		DueDate:            "15.05.2023.",
		MaxPoints:          10,
		Type:               models.First,
		Question:           FirstQuestion,
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
		Name:               "Homework 5",
		LastSubmissionDate: "15.05.2023.",
		DueDate:            "15.05.2023.",
		MaxPoints:          10,
		Type:               models.Second,
		Question:           SecondQuestion,
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
		Name:               "Homework 5",
		LastSubmissionDate: "09.05.2023.",
		DueDate:            "12.05.2023.",
		MaxPoints:          10,
		Type:               models.First,
		Question:           FirstQuestion,
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

func seedForum(storage Storage) {
	if err := storage.AddMessage(models.Message{
		Name:          "Ivo Ivić",
		Timestamp:     "8:25AM",
		MessageBody:   "Jeste li završili 3. zadaću?",
		IsCurrentUser: false,
	}); err != nil {
		panic(err)
	}

	if err := storage.AddMessage(models.Message{
		Name:          "Ana Anić",
		Timestamp:     "8:30AM",
		MessageBody:   "Ja jesam, dobila sam sve bodove",
		IsCurrentUser: false,
	}); err != nil {
		panic(err)
	}

	if err := storage.AddMessage(models.Message{
		Name:          "Marko Marković",
		Timestamp:     "8:45AM",
		MessageBody:   "Postoji sličan u auditornim vježbama",
		IsCurrentUser: true,
	}); err != nil {
		panic(err)
	}

	if err := storage.AddMessage(models.Message{
		Name:          "Pero Perić",
		Timestamp:     "8:47AM",
		MessageBody:   "Možeš li objasniti kako dobiti rješenje?",
		IsCurrentUser: false,
	}); err != nil {
		panic(err)
	}
}
