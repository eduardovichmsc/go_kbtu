package main

import (
	"log"
	"net/http"
	"time"

	"github.com/eduardovichmsc/practice_1/internal/handler"
)

const API_URL = "https://jsonplaceholder.typicode.com"
const HOST = "localhost:8080"
func main() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	postHandler := handler.NewPostHandler(client, API_URL)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /posts", postHandler.GetAllPosts)
	mux.HandleFunc("GET /posts/{id}", postHandler.GetPostByID)
	mux.HandleFunc("POST /posts", postHandler.CreatePost)
	
	log.Printf("\nStarted on %s\n", HOST)
	err := http.ListenAndServe(HOST, mux)
	if err != nil {
		log.Fatal(err)
	}
}