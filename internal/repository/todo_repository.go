package repository

import "database/sql"

type TodoRepository struct {
	queries *Queries
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{
		queries: New(db),
	}
}
