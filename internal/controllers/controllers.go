package controllers

import (
	"net/http"

	"github.com/LucasCarrias/ginurl/internal/shortener"
	"github.com/gin-gonic/gin"
)

type FormData struct {
	Url string `form:"url"`
}

const BaseUrl = "localhost:8080/"

var shortenedUrls = map[string]string{}

func RootHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func ShortenHandler(c *gin.Context) {
	data := &FormData{}
	c.Bind(data)

	shortenedUrl := shortener.BuildShortenedUrl(data.Url)

	shortenedUrls[shortenedUrl] = data.Url

	c.HTML(http.StatusOK, "url.html", gin.H{
		"url": BaseUrl + shortenedUrl,
	})
}

func RedirectHandler(c *gin.Context) {
	code := c.Param("code")

	url, exists := shortenedUrls[code]
	if exists {
		c.Redirect(http.StatusPermanentRedirect, url)
	} else {
		c.String(http.StatusNotFound, "Not Found")
	}
}