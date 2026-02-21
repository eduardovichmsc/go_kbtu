package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID     string
	Genre  string
	Budget int32
	Title  string
}

func main() {
	r := mux.NewRouter()

	movies := []Movie{}

	movies = append(movies, Movie{ID: "1", Genre: "horror", Budget: 500000, Title: "SAW"})
	movies = append(movies, Movie{ID: "1", Genre: "horror", Budget: 500000, Title: "SAW"})

	http.Handle("/", r)
}