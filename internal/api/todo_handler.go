package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/oliverslade/todo-api/internal/repository"
)

type TodoHandler struct {
	repo repository.TodoRepository
}

func NewTodoHandler(repo repository.TodoRepository) *TodoHandler {
	return &TodoHandler{repo: repo}
}

func (h *TodoHandler) ListTodo(w http.ResponseWriter, r *http.Request) {
	todos, err := h.repo.GetAllTodos(r.Context())
	if err != nil {
		slog.Error("failed to get todos", slog.String("error", err.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"failed to get todos"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		slog.Error("failed to parse id", slog.String("error", err.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"id must be an integer"}`))
		return
	}

	todo, err := h.repo.GetTodoById(r.Context(), id)
	if err != nil {
		slog.Error("failed to get todo", slog.String("error", err.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"failed to get todo"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)	
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var model = repository.Todo{}
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		slog.Error("failed to decode todo request", slog.String("error", err.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"failed to decode todo request"}`))
		return
	}

	if model.Message == "" {
		slog.Warn("message is required")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"message is required"}`))
		return
	}

	err = h.repo.CreateTodo(r.Context(), model)
	if err != nil {
		slog.Error("failed to create todo", slog.String("error", err.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"failed to create todo"}`))
		return
	}

	slog.Info("todo created successfully", slog.Int64("id", model.ID))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	returnNotImplemented(w)
}

func returnNotImplemented(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"error": "not implemented"})
}
