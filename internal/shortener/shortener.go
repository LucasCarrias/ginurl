package shortener

import "math/rand"

func BuildShortenedUrl(url string) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result []byte

	for i := 0; i < 5; i++ {
		result = append(result, chars[rand.Intn(len(chars))])
	}

	return string(result)
}