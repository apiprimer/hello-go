package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	From   string `json:"from"`
	Message string `json:"message"`
	Status bool `json:"status"`
}

type Request struct {
	Name   string `json:"name"`
}

func getHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		From:    "Go",
		Message: "Hello, Go!",
		Status:  true,
	})
}

// createUserHandler creates a new user
func postHello(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Simple validation (add more robust validation in a real app)
	if req.Name == "" {
		http.Error(w, "Name required", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{
		From:    "Go",
		Message: fmt.Sprintf("Hello, %s!", req.Name),
		Status:  true,
	})
}

func main() {
	mux := http.NewServeMux()

	// Register handlers
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getHello(w, r)
		case http.MethodPost:
			postHello(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server starting on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
