package repository

import "context"

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo Todo) error
	GetAllTodos(ctx context.Context) ([]Todo, error)
}
