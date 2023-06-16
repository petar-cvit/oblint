package main

import (
	"example.com/oblint/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	storage := internal.NewStorage()
	handler := internal.NewHandler(storage)

	r.GET("/test", handler.Test)

	r.Run()
}
