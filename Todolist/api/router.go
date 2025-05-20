package api

import (
	"github.com/go-chi/chi"
)

func (app *server) routes() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/todo", func(r chi.Router) {
		r.Mount("/", app.todoHandler.Routes())
	})

	return router
}
