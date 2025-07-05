package main

import (
	"fmt"
	"log"
	"net/http"

	"mood-service/controllers"
)

func main() {

	   // CORS middleware wrapper
	   corsHandler := func(h http.HandlerFunc) http.HandlerFunc {
			   return func(w http.ResponseWriter, r *http.Request) {
					   w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
					   w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
					   w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
					   if r.Method == "OPTIONS" {
							   w.WriteHeader(http.StatusOK)
							   return
					   }
					   h(w, r)
			   }
	   }

	   http.HandleFunc("/recommend", corsHandler(controllers.HandleRecommend))

	fmt.Println("Starting server on port 8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	fmt.Println("server is running on port 8081")


}