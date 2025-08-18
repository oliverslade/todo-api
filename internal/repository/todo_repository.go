package repository

import (
	"context"
	"database/sql"
)

type TodoRepo struct {
	queries *Queries
}

func NewTodoRepo(db *sql.DB) TodoRepository {
	return &TodoRepo{
		queries: New(db),
	}
}

func (r *TodoRepo) CreateTodo(ctx context.Context, todo Todo) error {
	return r.queries.CreateTodo(ctx, CreateTodoParams{
		Message:    todo.Message,
		IsFinished: todo.IsFinished,
	})
}

func (r *TodoRepo) GetAllTodos(ctx context.Context) ([]Todo, error) {
	return r.queries.GetAllTodos(ctx)
}
