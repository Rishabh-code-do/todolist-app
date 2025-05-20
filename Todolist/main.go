package main

import (
	"log"
	"todolist/api"
	"todolist/config"
	"todolist/connections"
	db "todolist/db/sqlc"
)

func main() {
	config, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	conn := connections.CreatePostgresSession(config.POSTGRES_CONNECTION)
	defer conn.Close()

	store := db.NewStore(conn)

	startHttpServer(api.ServerConfig{
		Store:  store,
		Config: config,
	})
}

func startHttpServer(arg api.ServerConfig) {
	app := api.NewServer(arg)
	log.Default().Println("Starting server on port", arg.Config.PORT)
	if err := app.Start(); err != nil {
		panic(err.Error())
	}
}
