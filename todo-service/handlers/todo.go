package handlers

import (
	"encoding/json"
	"github.com/D7NAMITE/todo-application.git/db"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ApplicationReponse struct {
	GOOD int
}

type TodoHandler struct {
	databaseUrl string
}

func NewTodoHandler(
	databaseUrl string,
) *TodoHandler {

	return &TodoHandler{
		databaseUrl: databaseUrl,
	}
}

func (h *TodoHandler) HandleGetTodoByClerkID(w http.ResponseWriter, r *http.Request) {
	todo_data, err := db.GetTodoByClerkID(chi.URLParam(r, "clerk_id"), h.databaseUrl)
	if err != nil {
		http.Error(w, "Failed to get Todos", http.StatusInternalServerError)
		return
	}

	todoJson, err := json.Marshal(todo_data)
	if err != nil {
		respondInternalServerError(w, err)
		return
	}

	w.Write(todoJson)
}

func (h *TodoHandler) HandlerCreateTodo(w http.ResponseWriter, r *http.Request) {

	err := db.CreateTodo(r, h.databaseUrl)

	if err != nil {
		respondInternalServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
