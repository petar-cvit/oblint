package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	storage Storage
}

func NewHandler(storage Storage) Handler {
	return Handler{
		storage: storage,
	}
}

func (h *Handler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, struct {
		Name string
	}{
		Name: "lasighas.rikjngh",
	})
}
