package main

import (
	"github.com/D7NAMITE/todo-application.git/config"
	"github.com/D7NAMITE/todo-application.git/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()
	router := chi.NewRouter()

	todoHandler := handlers.NewTodoHandler(
		cfg.DatabaseURL,
	)

	router.Route("/todo", func(r chi.Router) {
		r.Get("/user/{user_id}", todoHandler.HandleGetTodoByUserid)
		r.Post("/new", todoHandler.HandlerCreateTodo)
	})

	http.ListenAndServe(":8080", router)
}
