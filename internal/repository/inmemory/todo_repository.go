package inmemory

import (
	"context"

	"github.com/oliverslade/todo-api/internal/repository"
)

type InMemoryTodoRepo struct {
	todos []repository.Todo
}

func NewInMemoryTodoRepo() repository.TodoRepository {
	return &InMemoryTodoRepo{
		todos: []repository.Todo{
			{
				ID:         1,
				Message:    "Buy groceries",
				IsFinished: false,
			},
		},
	}
}

func (r *InMemoryTodoRepo) CreateTodo(ctx context.Context, todo repository.Todo) error {
	r.todos = append(r.todos, todo)
	return nil
}

func (r *InMemoryTodoRepo) GetAllTodos(ctx context.Context) ([]repository.Todo, error) {
	return r.todos, nil
}
