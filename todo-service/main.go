package main

import (
	"database/sql"
	"fmt"
	"github.com/D7NAMITE/todo-application.git/config"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()
	router := chi.NewRouter()

	todoHandler := handlers.todo
}
