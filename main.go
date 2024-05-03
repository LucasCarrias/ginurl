package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")


	router.GET("/", rootHandler)

	router.POST("/shorten", shortenHandler)
	router.Run(":8080")
}

func rootHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

type FormData struct {
	Url string `form:"url"`
}

func shortenHandler(c *gin.Context) {
	data := &FormData{}

	c.Bind(data)
	c.HTML(http.StatusOK, "url.html", gin.H{
		"title": data.Url,
	})
}