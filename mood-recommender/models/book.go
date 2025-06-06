package models


type Book struct {
	Title    string   `json:"title"`
	Author   string   `json:"author"`
	Genre    string   `json:"genre"`
	MoodTags []string `json:"mood_tags"`
}

