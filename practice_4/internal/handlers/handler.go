package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Movie struct {
	ID     string
	Genre  string
	Budget int32
	Title  string
}

// GET /tasks
// GET /tasks?id=1
func (h *TaskHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		json.NewEncoder(w).Encode(h.Tasks)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid id"})
		return
	}

	for _, task := range h.Tasks {
		if task.ID == id {
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "task not found"})
}