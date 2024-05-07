package controllers

import (
	"net/http"

	"github.com/LucasCarrias/ginurl/internal/models"
	"github.com/LucasCarrias/ginurl/internal/shortener"
	"github.com/gin-gonic/gin"
)

type FormData struct {
	Url string `form:"url"`
}

func RootHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func ShortenHandler(c *gin.Context) {
	data := &FormData{}
	c.Bind(data)

	code := shortener.CreateCode()

	url := models.CreateShortenedUrl(code, data.Url)

	c.HTML(http.StatusOK, "url.html", gin.H{
		"url": url.UrlWithCode(),
	})
}

func RedirectHandler(c *gin.Context) {
	code := c.Param("code")

	err, result := models.GetShortenedUrl(code)

	if err != nil {
		c.String(http.StatusNotFound, "Not Found")
	} else {
		c.Redirect(http.StatusPermanentRedirect, result.Source)
	}
}