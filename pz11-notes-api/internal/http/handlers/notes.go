package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/notes-api/internal/core/service"
	"example.com/notes-api/internal/repo"
	"github.com/go-chi/chi/v5"
	"github.com/swaggo/swag"
)

type Handler struct {
	Service *service.NoteService
}

type NoteCreate struct {
	Title   string `json:"title" example:"Новая заметка"`
	Content string `json:"content" example:"Текст заметки"`
}

type NoteUpdate struct {
	Title   string `json:"title,omitempty" example:"Обновлено"`
	Content string `json:"content,omitempty" example:"Новый текст"`
}

func NewHandler(repo *repo.NoteRepoMem) *Handler {
	service := service.NewNoteService(repo)
	return &Handler{Service: service}
}

// GetSwaggerSpec возвращает JSON спецификацию для ReDoc
func GetSwaggerSpec() []byte {
	spec, _ := swag.ReadDoc()
	return []byte(spec)
}

// CreateNote godoc
// @Summary      Создать заметку
// @Description  Создание новой заметки (требуется авторизация)
// @Tags         notes
// @Accept       json
// @Produce      json
// @Param        input  body     NoteCreate  true  "Данные новой заметки"
// @Success      201    {object} core.Note
// @Failure      400    {object} map[string]string
// @Failure      401    {object} map[string]string
// @Failure      500    {object} map[string]string
// @Security     BearerAuth
// @Router       /notes [post]
func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	if token := r.Header.Get("Authorization"); token == "" {
		http.Error(w, "Authorization required", http.StatusUnauthorized)
		return
	}
	
	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	note, err := h.Service.CreateNote(input.Title, input.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

// List godoc
// @Summary      Список заметок
// @Description  Возвращает список заметок (публичный доступ)
// @Tags         notes
// @Success      200    {array}  core.Note
// @Failure      500    {object} map[string]string
// @Router       /notes [get]
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	notes, err := h.Service.GetAllNotes()
	if err != nil {
		http.Error(w, "Failed to retrieve notes", http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

// GetNote godoc
// @Summary      Получить заметку
// @Description  Получение заметки по ID (требуется авторизация)
// @Tags         notes
// @Param        id   path   int  true  "ID"
// @Success      200  {object}  core.Note
// @Failure      404  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Security     BearerAuth
// @Router       /notes/{id} [get]
func (h *Handler) GetNote(w http.ResponseWriter, r *http.Request) {
	if token := r.Header.Get("Authorization"); token == "" {
		http.Error(w, "Authorization required", http.StatusUnauthorized)
		return
	}
	
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}
	
	note, err := h.Service.GetNoteByID(id)
	if err != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

// EditNote godoc
// @Summary      Обновить заметку (частично)
// @Description  Частичное обновление заметки (требуется авторизация)
// @Tags         notes
// @Accept       json
// @Param        id     path   int        true  "ID"
// @Param        input  body   NoteUpdate true  "Поля для обновления"
// @Success      200    {object}  core.Note
// @Failure      400    {object}  map[string]string
// @Failure      401    {object}  map[string]string
// @Failure      404    {object}  map[string]string
// @Security     BearerAuth
// @Router       /notes/{id} [patch]
func (h *Handler) EditNote(w http.ResponseWriter, r *http.Request) {
	if token := r.Header.Get("Authorization"); token == "" {
		http.Error(w, "Authorization required", http.StatusUnauthorized)
		return
	}
	
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}
	
	var input NoteUpdate
	
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	
	note, err := h.Service.UpdateNote(id, input.Title, input.Content)
	if err != nil {
		if err.Error() == "note not found" {
			http.Error(w, "Note not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

// DeleteNote godoc
// @Summary      Удалить заметку
// @Description  Удаление заметки по ID (требуется авторизация)
// @Tags         notes
// @Param        id  path  int  true  "ID"
// @Success      204  "No Content"
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Security     BearerAuth
// @Router       /notes/{id} [delete]
func (h *Handler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	// Пример проверки авторизации
	if token := r.Header.Get("Authorization"); token == "" {
		http.Error(w, "Authorization required", http.StatusUnauthorized)
		return
	}
	
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}
	
	err = h.Service.DeleteNote(id)
	if err != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}