package repository

import "database/sql"

type TodoRepository struct {
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{}
}
