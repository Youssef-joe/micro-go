package controllers

import (
	"encoding/json"
	"net/http"
	"os"

	"mood-service/utils"
	"mood-service/models"
)




func HandleRecommend(w http.ResponseWriter, r *http.Request) {
	mood := r.URL.Query().Get("mood")
	if mood == "" {
		http.Error(w, "Mood parameter is required", http.StatusBadRequest)
		return
	}

	file, err := os.ReadFile("data/books.json")
	if err != nil {
		http.Error(w, "Failed to read data", http.StatusInternalServerError)
		return
	}

	var books []models.Book
	if err := json.Unmarshal(file, &books); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusInternalServerError)
		return
	}

	matched := utils.MatchBooksByMood(mood, books)

	response := map[string]interface{}{
		"mood":  mood,
		"books": matched,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
