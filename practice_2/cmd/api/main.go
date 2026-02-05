package main

import (
	"log"
	"net/http"
	"todo/internal/handlers"
	"todo/internal/middleware"
)

func main() {
	HOST := "localhost:8080"
	
	taskHandler := handlers.NewTaskHandler()
	mux := http.NewServeMux()

	handler := middleware.LoggingMiddleware(
		middleware.AuthMiddleware(taskHandler),
	)

	mux.Handle("/tasks", handler)

	log.Printf("\nStarted on %s\n", HOST);
	err := http.ListenAndServe(HOST, mux)
	if err != nil {
		log.Fatal(err)
	}
}