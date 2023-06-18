package internal

import "example.com/oblint/internal/models"

func seedHistory(storage Storage) {
	err := storage.SaveToHistory(models.HistoryHomework{
		ID:             "123",
		Name:           "Prva zadaca",
		SubmissionDate: "12.5.2023.",
		DueDate:        "11.5.2023.",
		Points:         5,
		MaxPoints:      10,
		Type:           models.First,
		Correct:        "0,1,1,0",
		Answered:       "0,1,0,1",
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
		Correct:        "A or B and not A",
		Answered:       "A or B and not B",
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
		Correct:        "0,0,1,0",
		Answered:       "0,1,0,1",
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
		Answered:           "",
		Started:            false,
	})
	if err != nil {
		panic(err)
	}

	err = storage.SaveToHomeworks(models.Homework{
		ID:                 "322",
		Name:               "Prva zadaca za napraviti",
		LastSubmissionDate: "",
		DueDate:            "15.5.2023.",
		MaxPoints:          8,
		Type:               models.Second,
		Answered:           "",
		Started:            false,
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
		Answered:           "1,_,0,1",
		Started:            true,
	})
	if err != nil {
		panic(err)
	}
}
