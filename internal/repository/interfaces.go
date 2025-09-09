package repository

import "context"

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo Todo) error
	GetAllTodos(ctx context.Context) ([]Todo, error)
	GetTodoById(ctx context.Context, id int64) (Todo, error)
	SetTodoFinished(ctx context.Context, id int64, isFinished bool) error
}
