package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	redisHost := os.Getenv("REDIS_HOST")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	rc, err := NewRedisClient(redisHost, redisPassword)

	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Ready to take requests...")

	defer rc.Close()

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			username := r.URL.Query().Get("username")
			log.Println("searching for user ", username)
			user, err := rc.GetUser(username)
			if err != nil {
				log.Println("Error while searching for user ", username)
				http.Error(w, fmt.Sprintf("Failed to get user: %v", err), http.StatusInternalServerError)
				return
			}

			if user == nil {
				log.Println("Error!! user does not exist: ", username)
				http.NotFound(w, r)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
		case http.MethodPost:
			var user User

			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				http.Error(w, fmt.Sprintf("Failed to decode user: %v", err), http.StatusBadRequest)
				return
			}

			if err := rc.SetUser(user); err != nil {
				log.Println("Error while storing user ", user.Username)
				http.Error(w, fmt.Sprintf("Failed to set user: %v", err), http.StatusInternalServerError)
				return
			}

			fmt.Println("Successfully created user ", user.Username)

			w.WriteHeader(http.StatusCreated)
		case http.MethodDelete:
			username := r.URL.Query().Get("username")
			log.Println("Trying to delete user ", username)

			if err := rc.DeleteUser(username); err != nil {
				log.Println("Error while deleting user ", username)
				http.Error(w, fmt.Sprintf("Failed to delete user: %v", err), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusNoContent)
		default:
			http.Error(w, "Invalid method", http.StatusBadRequest)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
