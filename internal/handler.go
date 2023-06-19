package internal

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"time"

	"example.com/oblint/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Handler struct {
	storage Storage
}

func NewHandler(storage Storage) Handler {
	// seed
	storage.Clear()
	seedHistory(storage)
	seedOngoing(storage)
	seedForum(storage)

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

func (h *Handler) Homeworks(c *gin.Context) {
	history, err := h.storage.GetHomeworks()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, history)
}

func (h *Handler) HomeworkById(c *gin.Context) {
	ID := c.Param("ID")

	history, err := h.storage.GetHomeworkByID(ID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, history)
}

func (h *Handler) Stats(c *gin.Context) {
	generationScore := []int{10, 3, 8, 6, 3, 6, 8, 3, 9, 7, 0, 7, 9, 8, 3, 8, 9, 4, 10, 5}

	minGeneration := generationScore[0]
	maxGeneration := generationScore[0]
	for _, score := range generationScore {
		if score < minGeneration {
			minGeneration = score
		}
		if score > maxGeneration {
			maxGeneration = score
		}
	}

	history, err := h.storage.GetHistory()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	for _, homework := range history {
		generationScore = append(generationScore, homework.Points)
	}

	ongoing, err := h.storage.GetHomeworks()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	inProgress := 0
	notStarted := 0
	for _, homework := range ongoing {
		if homework.Started {
			inProgress++
		} else {
			notStarted++
		}
	}

	stats := models.Stats{
		GenerationScore: generationScore,
		Min:             float64(minGeneration),
		Max:             float64(maxGeneration),
		Avg:             calculateAverage(generationScore),
		Std:             calculateStandardDeviation(generationScore),
		Finished:        len(history),
		InProgress:      inProgress,
		NotStarted:      notStarted,
	}

	c.JSON(http.StatusOK, stats)
}

func (h *Handler) SubmitHomework(c *gin.Context) {
	var req models.SubmitHomeworkDTO
	if err := c.BindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if req.DataAnswer == nil {
		req.DataAnswer = []string{}
	}

	hw, err := h.storage.GetHomeworkByID(req.ID)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			c.String(http.StatusBadRequest, "homework not found")
			return
		}

		c.Status(http.StatusBadRequest)
		return
	}

	if hw.Type == 2 && len(req.DataAnswer) != 12 {
		c.String(http.StatusBadRequest, "answer should be long 12")
		return
	}

	if err := h.storage.SaveToHistory(models.HistoryHomework{
		ID:             hw.ID,
		Name:           hw.Name,
		SubmissionDate: time.Now().Format("02.01.2006."),
		DueDate:        hw.DueDate,
		Points:         0,
		MaxPoints:      hw.MaxPoints,
		Type:           hw.Type,
		Question:       hw.Question,
		Statement:      hw.Statement,
		Data:           req.DataAnswer,
		CorrectData:    hw.CorrectData,
		Answer:         req.StatementAnswer,
		CorrectAnswer:  hw.CorrectAnswer,
	}); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	h.storage.DeleteFromHomeworks(hw)

	c.Status(http.StatusOK)
}

func (h *Handler) GetMessages(c *gin.Context) {
	msgs, err := h.storage.GetForum()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, msgs)
}

func (h *Handler) SubmitMessage(c *gin.Context) {
	var msg models.Message
	if err := c.BindJSON(&msg); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	msg.Timestamp = time.Now().Format(time.Kitchen)

	if err := h.storage.AddMessage(msg); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
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

func calculateAverage(arr []int) float64 {
	sum := 0
	for _, num := range arr {
		sum += num
	}

	return math.Round((float64(sum)/float64(len(arr)))*100) / 100
}

func calculateStandardDeviation(arr []int) float64 {
	avg := calculateAverage(arr)
	variance := 0.0

	for _, num := range arr {
		variance += math.Pow(float64(num)-avg, 2)
	}

	variance /= float64(len(arr))
	stdDev := math.Sqrt(variance)
	return math.Round(stdDev*100) / 100
}
