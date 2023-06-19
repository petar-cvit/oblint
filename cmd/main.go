package main

import (
	"net/http"

	"example.com/oblint/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	storage, err := internal.NewStorage()
	if err != nil {
		panic(err)
	}

	handler := internal.NewHandler(storage)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "PING")
	})

	r.GET("/test", handler.Test)

	r.GET("/history", handler.HomeworkHistory)
	r.GET("/history/:ID", handler.HomeworkHistoryById)

	r.GET("/homeworks", handler.Homeworks)
	r.GET("/homeworks/:ID", handler.HomeworkById)
	r.POST("/homeworks/new", handler.CreateHomework)

	r.GET("/stats", handler.Stats)

	r.POST("/submit", handler.SubmitHomework)

	r.GET("/forum", handler.GetMessages)
	r.POST("/forum", handler.SubmitMessage)

	r.Run()
}
