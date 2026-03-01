package main

import (
	"fmt"
	"log"
	"net/http"

	"practice_5/internal/database"
	"practice_5/internal/handlers"
	"practice_5/internal/models"

	"github.com/gorilla/mux"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	defer db.Close()

	movieModel := &models.MovieModel{DB: db}
	movieHandler := &handlers.MovieHandler{Model: movieModel}

	r := mux.NewRouter()
	r.HandleFunc("/movies", movieHandler.GetMovies).Methods("GET")
	r.HandleFunc("/movies", movieHandler.CreateMovie).Methods("POST")
	
	r.HandleFunc("/movies/{id}", movieHandler.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", movieHandler.DeleteMovie).Methods("DELETE")

	fmt.Println("Starting the Server on :8000...")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}