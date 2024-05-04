package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

const BaseUrl = "localhost:8080/"

type FormData struct {
	Url string `form:"url"`
}

var shortenedUrls = map[string]string{}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", rootHandler)
	router.POST("/shorten", shortenHandler)
	router.GET("/:code", redirectHandler)
	router.Run(":8080")
}

func rootHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func shortenHandler(c *gin.Context) {
	data := &FormData{}
	c.Bind(data)

	shortenedUrl := BuildShortenedUrl(data.Url)

	shortenedUrls[shortenedUrl] = data.Url

	c.HTML(http.StatusOK, "url.html", gin.H{
		"url": BaseUrl + shortenedUrl,
	})
}

func redirectHandler(c *gin.Context) {
	code := c.Param("code")

	url, exists := shortenedUrls[code]
	if exists {
		c.Redirect(http.StatusPermanentRedirect, url)
	} else {
		c.String(http.StatusNotFound, "Not Found")
	}
}

func BuildShortenedUrl(url string) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result []byte

	for i := 0; i < 5; i++ {
		result = append(result, chars[rand.Intn(len(chars))])
	}

	return string(result)
}
