package utils


import (
	"strings"

	"mood-service/models"
)

func MatchBooksByMood(mood string, books []models.Book) []models.Book {
	var matched []models.Book
	for _, book := range books {
		for _, tag := range book.MoodTags {
			if strings.Contains(strings.ToLower(tag), strings.ToLower(mood)) {
				matched = append(matched, book)
				break
			}
		}
	}
	return matched
}