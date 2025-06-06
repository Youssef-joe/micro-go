package main
import (
	"fmt"
	"log"
	"net/http"

	"mood-service/controllers"
)

func main() {
	http.HandleFunc("/recommend", controllers.HandleRecommend)

	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}