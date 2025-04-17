package handlers

import (
	"encoding/json"
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

func (h *TodoHandler) HandlGetTodoByUserid(w http.ResponseWriter, r *http.Request) {

	todo_data, err := db.SearchProjectByParameterQuery(r, h.databaseUrl)

	if err != nil {
		http.Error(w, "Failed to get project", http.StatusInternalServerError)
		return
	}

	repoJson, err := json.Marshal(todo_data)

	if err != nil {
		respondInternalServerError(w, err)
		return
	}

	w.Write(repoJson)
}
