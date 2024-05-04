package main

import (
	"github.com/LucasCarrias/ginurl/internal/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", controllers.RootHandler)
	router.POST("/shorten", controllers.ShortenHandler)
	router.GET("/:code", controllers.RedirectHandler)
	router.Run(":8080")
}

