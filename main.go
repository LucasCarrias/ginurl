package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

const BaseUrl = "localhost:8080/"

type ShortenedUrl struct {
	original string
	url string
}

type FormData struct {
	Url string `form:"url"`
}

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

func shortenHandler(c *gin.Context) {
	data := &FormData{}
	c.Bind(data)

	shortenedUrl := BuildShortenedUrl(data.Url)

	c.HTML(http.StatusOK, "url.html", gin.H{
		"url": shortenedUrl.url,
	})
}

func BuildShortenedUrl(url string) ShortenedUrl {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result []byte

	for i := 0; i < 5; i++ {
		result = append(result, chars[rand.Intn(len(chars))])
	}

	return ShortenedUrl{
		original: url,
		url: BaseUrl + string(result),
	}
}
