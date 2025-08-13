package api

import (
	"encoding/json"
	"net/http"

	"github.com/oliverslade/todo-api/internal/repository"
)

type TodoHandler struct {
	repo *repository.TodoRepository
}

func NewTodoHandler(repo *repository.TodoRepository) *TodoHandler {
	return &TodoHandler{repo: repo}
}

func (h *TodoHandler) ListTodo(w http.ResponseWriter, r *http.Request) {
	returnNotImplemented(w)
}

func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	returnNotImplemented(w)
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	returnNotImplemented(w)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	returnNotImplemented(w)
}

func returnNotImplemented(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"error": "not implemented"})
}
