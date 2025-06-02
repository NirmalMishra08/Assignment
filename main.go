package main

import (
	"assignment/db"
	sqlc "assignment/db/sqlc"
	"assignment/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	database, err := db.ConnectToDB()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer database.Close()
	queries := sqlc.New(database)
	if queries == nil {
		log.Fatal("Failed to create database queries")
	}

	h := handlers.NewHandler(queries)

	r := chi.NewRouter()

	h.RegisterRoutes(r)

	fmt.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed to start:", err)
	}
	log.Println("Server started on :8080")

}
