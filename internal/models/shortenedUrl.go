package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const BaseUrl = "localhost:8080/"

type ShortenedUrl struct {
	gorm.Model
	Code string `gorm:"unique"`
	Source string
}

func init() {
	db := db()

	db.AutoMigrate(&ShortenedUrl{})
}

func db() *gorm.DB  {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
    panic("failed to connect database")
  }

	return db
}

func CreateShortenedUrl(code, source string) ShortenedUrl {
	db := db()

	url := &ShortenedUrl{Code: code, Source: source}

	result := db.Create(url)

	if result.Error != nil {
		panic(result.Error)
	}

	return *url
}

func GetShortenedUrl(code string) (error, ShortenedUrl) {
	var url ShortenedUrl
	
	db := db()
	result := db.Where("code = ?", code).First(&url)

	return result.Error, url
}

func (s *ShortenedUrl) UrlWithCode() string {
	return BaseUrl + s.Code
}