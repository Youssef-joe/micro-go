package main

import (
	"encoding/json"
	"os"
	"log"
	"net/http"
	"sync"
)

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

var (
	users []User
	mu sync.Mutex
	
)

func loadUsers() {
	data, err := os.ReadFile("users.json")
	if err != nil {
		users = []User{}
		return
	}
	json.Unmarshal(data, &users)
}

func saveUsers() {
	data, _ := json.Marshal(users)
	os.WriteFile("users.json", data, 0644)
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err!=nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	newUser.ID = len(users) + 1
	users = append(users,newUser)
	saveUsers()
	
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func main() {
	loadUsers()


	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getUsers(w,r)
			} else if r.Method == http.MethodPost {
				addUser(w,r)
			} else {
				http.Error(w, "method are not allowed", http.StatusMethodNotAllowed)
			}
	})

	log.Println("user service running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}