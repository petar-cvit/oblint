package internal

import (
	"fmt"
	"net/http"

	"example.com/oblint/internal/models"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	storage Storage
}

func NewHandler(storage Storage) Handler {
	// seed
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

	return Handler{
		storage: storage,
	}
}

func (h *Handler) HomeworkHistory(c *gin.Context) {
	history, err := h.storage.GetHistory()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, history)
}

func (h *Handler) HomeworkHistoryById(c *gin.Context) {
	ID := c.Param("ID")

	history, err := h.storage.GetHistoryByID(ID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, history)
}

func (h *Handler) Test(c *gin.Context) {
	err := h.storage.Random()
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, struct {
		Name string
	}{
		Name: "lasighas.rikjngh",
	})
}
