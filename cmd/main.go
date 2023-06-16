package main

import (
	"net/http"

	"example.com/oblint/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	storage := internal.NewStorage()
	handler := internal.NewHandler(storage)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "PING")
	})

	r.GET("/test", handler.Test)

	r.Run()
}
