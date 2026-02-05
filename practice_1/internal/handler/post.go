package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/eduardovichmsc/practice_1/internal/models"
)

type PostHandler struct {
	Client *http.Client
	BaseURL string
}

func NewPostHandler(client *http.Client, baseURL string) *PostHandler {
	return &PostHandler{
		Client: client,
		BaseURL: baseURL,
	}
}

// GET /posts
func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request){
	res, err := h.Client.Get(h.BaseURL + "/posts")
	if err != nil {
		http.Error(w, "Error connecting to " + h.BaseURL, http.StatusBadGateway)
	}
	w.Header().Set("Content-Type", "applcation/json")
	w.WriteHeader(res.StatusCode)
	io.Copy(w, res.Body)
}

// GET /posts/:id
func (h *PostHandler) GetPostByID(w http.ResponseWriter, r *http.Request){
	id := r.PathValue("id")

	res, err := h.Client.Get(fmt.Sprintf("%s/posts/%s", h.BaseURL, id))
	if err != nil {
		http.Error(w, "Error: GetPostByID" + h.BaseURL, http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "applcation/json")
	w.WriteHeader(res.StatusCode)
	io.Copy(w, res.Body)
}

// POST /posts
func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(post)
	if err != nil {
		http.Error(w, "Внутренняя ошибка", http.StatusInternalServerError)
		return
	}

	resp, err := h.Client.Post(h.BaseURL+"/posts", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, "Ошибка при создании поста во внешнем API", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}