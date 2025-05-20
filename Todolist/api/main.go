package api

import (
	"fmt"
	"net/http"
	"todolist/api/todo"
	"todolist/config"
	db "todolist/db/sqlc"

	"github.com/go-chi/chi"
)

type ServerConfig struct {
	Store  db.Store
	Config *config.Config
}

type server struct {
	store       db.Store
	config      *config.Config
	todoHandler *todo.Handler
	router      *chi.Mux
}

func NewServer(arg ServerConfig) *server {
	app := server{
		config: arg.Config,
		store:  arg.Store,
	}

	app.todoHandler = todo.NewHandler(&todo.HandlerConfig{
		Config: arg.Config,
		Store:  arg.Store,
	})
	app.router = app.routes()
	return &app
}

func (app *server) Start() error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.config.PORT),
		Handler: app.router,
	}
	return srv.ListenAndServe()
}
