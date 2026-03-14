package handlers

import (
	"encoding/json"
	"net/http"

	"practice_6/internal/models"
)

type UserHandler struct {
	Model *models.User
}

func (h *UserHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := 
	h.Model.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if movies == nil {
		movies = []models.Movie{}
	}
	json.NewEncoder(w).Encode(movies)
}

func (h *UserHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var m models.Movie
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if err := h.Model.Insert(&m); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(m)
}

func (h *UserHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}