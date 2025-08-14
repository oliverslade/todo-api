package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/oliverslade/todo-api/internal/repository/inmemory"
	"github.com/stretchr/testify/assert"
)

func TestTodoHandler_CreateTodo(t *testing.T) {

	t.Run("should create a todo successfully with valid request", func(t *testing.T) {
		repo := inmemory.NewInMemoryTodoRepo()
		handler := NewTodoHandler(repo)

		router := chi.NewRouter()
		router.Post("/todos", handler.CreateTodo)

		request := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewBufferString(`{"message": "test", "is_finished": false}`))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)
		assert.Equal(t, http.StatusCreated, response.Code)
	})

	t.Run("should return bad request when unable to decode request", func(t *testing.T) {
		repo := inmemory.NewInMemoryTodoRepo()
		handler := NewTodoHandler(repo)

		router := chi.NewRouter()
		router.Post("/todos", handler.CreateTodo)

		request := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewBufferString(`{:D:BLAH"is_finished": false}`))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "{\"error\":\"failed to decode todo request\"}", response.Body.String())
	})

	t.Run("should return bad requesr when message is empty", func(t *testing.T) {
		repo := inmemory.NewInMemoryTodoRepo()
		handler := NewTodoHandler(repo)

		router := chi.NewRouter()
		router.Post("/todos", handler.CreateTodo)

		request := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewBufferString(`{"message": "", "is_finished": false}`))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "{\"error\":\"message is required\"}", response.Body.String())
	})
}
