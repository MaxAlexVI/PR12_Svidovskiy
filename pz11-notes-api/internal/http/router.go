package httpx

import (
	"embed"
	"net/http"

	"example.com/notes-api/internal/http/handlers"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

//go:embed redoc.html
var redocTemplate embed.FS

func NewRouter(h *handlers.Handler) *chi.Mux {
	r := chi.NewRouter()
  
	// Swagger UI
	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("doc.json"),
	))

	// ReDoc (альтернативная документация)
	r.Get("/redoc", func(w http.ResponseWriter, r *http.Request) {
		data, err := redocTemplate.ReadFile("redoc.html")
		if err != nil {
			http.Error(w, "Template not found", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write(data)
	})

	// Swagger JSON для ReDoc
	r.Get("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		// Используем встроенную спецификацию
		swaggerSpec := handlers.GetSwaggerSpec()
		w.Header().Set("Content-Type", "application/json")
		w.Write(swaggerSpec)
	})

	// API routes
	r.Post("/api/v1/notes", h.CreateNote)
	r.Get("/api/v1/notes", h.List)
	r.Get("/api/v1/notes/{id}", h.GetNote)
	r.Patch("/api/v1/notes/{id}", h.EditNote)
	r.Delete("/api/v1/notes/{id}", h.DeleteNote)
	
	return r
}