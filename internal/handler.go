package internal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	storage Storage
}

func NewHandler(storage Storage) Handler {
	// seed
	seedHistory(storage)
	seedOngoing(storage)

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
