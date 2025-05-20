package todo

import (
	"todolist/config"
	db "todolist/db/sqlc"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type Handler struct {
	config *config.Config
	store  db.Store
}

type HandlerConfig struct {
	Config *config.Config
	Store  db.Store
}

func NewHandler(config *HandlerConfig) *Handler {
	return &Handler{
		config: config.Config,
		store:  config.Store,
	}
}

func (h *Handler) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	router.Post("/create", h.createTodo)
	router.Get("/get", h.getTodos)
	router.Get("/get/{id}", h.getTodosById)
	router.Patch("/update/{id}", h.updateTodo)
	router.Delete("/delete/{id}", h.deleteTodo)
	router.Patch("/complete/{id}", h.markTodoCompleted)
	return router
}
