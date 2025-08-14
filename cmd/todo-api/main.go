package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/oliverslade/todo-api/internal/api"
	"github.com/oliverslade/todo-api/internal/repository"
)

func main() {

	db, err := sql.Open("sqlite3", "todo.db")
	if err != nil {
		slog.Error("failed to open db")
		os.Exit(1)
	}

	repository := repository.NewTodoRepository(db)
	todoHandler := api.NewTodoHandler(repository)

	router := chi.NewRouter()
	router.Get("/todos", todoHandler.ListTodo)
	router.Get("/todos/{id}", todoHandler.GetTodo)
	router.Post("/todos", todoHandler.CreateTodo)
	router.Put("/todos/{id}", todoHandler.UpdateTodo)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()

}
