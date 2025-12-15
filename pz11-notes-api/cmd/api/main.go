// Package main Notes API server.
//
// @title           Notes API
// @version         1.0
// @description     Учебный REST API для заметок (CRUD).
// @contact.name    Backend Course
// @contact.email   example@university.ru
// @BasePath        /api/v1
package main

import (
	"log"
	"net/http"

	_ "example.com/notes-api/docs"
	httpx "example.com/notes-api/internal/http"
	"example.com/notes-api/internal/http/handlers"
	"example.com/notes-api/internal/repo"
)

func main() {
	repo := repo.NewNoteRepoMem()
	h := handlers.NewHandler(repo) 
	r := httpx.NewRouter(h)

	log.Println("Server started at :8080")
	log.Println("Swagger UI: http://localhost:8080/swagger/index.html")
	log.Fatal(http.ListenAndServe(":8080", r))
}