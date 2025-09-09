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

func TestTodoHandler_ListTodo(t *testing.T) {

	t.Run("should return all todos successfully with valid request", func(t *testing.T) {
		repo := inmemory.NewInMemoryTodoRepo()
		handler := NewTodoHandler(repo)

		router := chi.NewRouter()
		router.Get("/todos", handler.ListTodo)

		request := httptest.NewRequest(http.MethodGet, "/todos", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, `[{"id":1,"message":"Buy groceries","is_finished":false}]`+"\n", response.Body.String())
	})
}

func TestTodoHandler_GetTodoById(t *testing.T) {

	t.Run("should return a todo successfully with valid request", func(t *testing.T) {
		repo := inmemory.NewInMemoryTodoRepo()
		handler := NewTodoHandler(repo)

		router := chi.NewRouter()
		router.Get("/todos/{id}", handler.GetTodo)

		request := httptest.NewRequest(http.MethodGet, "/todos/1", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, `{"id":1,"message":"Buy groceries","is_finished":false}`+"\n", response.Body.String())
	})

	t.Run("should return bad request when id is not an integer", func(t *testing.T) {
		repo := inmemory.NewInMemoryTodoRepo()
		handler := NewTodoHandler(repo)

		router := chi.NewRouter()
		router.Get("/todos/{id}", handler.GetTodo)

		request := httptest.NewRequest(http.MethodGet, "/todos/not-an-integer", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "{\"error\":\"id must be an integer\"}", response.Body.String())
	})
}

func TestTodoHandler_UpdateTodo(t *testing.T) {

	t.Run("should update a todo successfully with only is_finished", func(t *testing.T) {
		repo := inmemory.NewInMemoryTodoRepo()
		handler := NewTodoHandler(repo)

		router := chi.NewRouter()
		router.Put("/todos/{id}", handler.UpdateTodo)

		request := httptest.NewRequest(http.MethodPut, "/todos/1", bytes.NewBufferString(`{"is_finished": true}`))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, `{"id":1,"message":"Buy groceries","is_finished":true}`+"\n", response.Body.String())
	})
}
