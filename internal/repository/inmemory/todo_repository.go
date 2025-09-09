package inmemory

import (
	"context"
	"errors"

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

func (r *InMemoryTodoRepo) GetTodoById(ctx context.Context, id int64) (repository.Todo, error) {
	for _, todo := range r.todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return repository.Todo{}, errors.New("todo not found")
}

func (r *InMemoryTodoRepo) SetTodoFinished(ctx context.Context, id int64, isFinished bool) error {
	for i := range r.todos {
		if r.todos[i].ID == id {
			r.todos[i].IsFinished = isFinished
			return nil
		}
	}
	return errors.New("todo not found")
}
