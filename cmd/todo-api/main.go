package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oliverslade/todo-api/internal/api"
	"github.com/oliverslade/todo-api/internal/repository"
)

func main() {

	repository := repository.NewTodoRepository()
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
