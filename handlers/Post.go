package handlers

import (
	db "assignment/db/sqlc"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var req db.CreatePostParams
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	post, err := h.queries.CreatePost(r.Context(), req)
	if err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		fmt.Println("CreatePost error:", err)
		return
	}

	json.NewEncoder(w).Encode(post)

}

func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {

	post, err := h.queries.ListPosts(r.Context())
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		fmt.Println("GetPost error:", err)
		return
	}

	json.NewEncoder(w).Encode(post)
}

func (h *Handler) GetPostById(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		fmt.Println("GetUserById error:", err)
		return
	}

	post, err := h.queries.GetPost(r.Context(), int32(userIdInt))
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		fmt.Println("GetUserById error:", err)
		return
	}

	json.NewEncoder(w).Encode(post)
}
func (h *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		fmt.Println("GetUserById error:", err)
		return
	}

	err = h.queries.DeletePost(r.Context(), int32(userIdInt))
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		fmt.Println("GetUserById error:", err)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Post deleted successfully"})

}

func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		fmt.Println("GetUserById error:", err)
		return
	}

	var req db.UpdatePostParams
	req.ID = int32(userIdInt)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	post, err := h.queries.UpdatePost(r.Context(), req)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		fmt.Println("GetUserById error:", err)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"post":    post,
		"message": "Post updated successfully",
	})
}
