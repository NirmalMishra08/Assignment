package handlers

import (
	db "assignment/db/sqlc"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	queries *db.Queries
}

func NewHandler(q *db.Queries) *Handler {
	return &Handler{queries: q}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/posts", h.CreatePost)
		r.Get("/posts", h.GetPost)
		r.Get("/post/{id}", h.GetPostById)
		r.Put("/post/{id}", h.UpdatePost)
		r.Delete("/post/{id}", h.DeletePost)

	})
}
