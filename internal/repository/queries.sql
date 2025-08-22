-- name: GetAllTodos :many
SELECT id, message, is_finished FROM todos
ORDER BY id;

-- name: GetTodoById :one
SELECT id, message, is_finished FROM todos
WHERE id = ?;

-- name: CreateTodo :exec
INSERT INTO todos (message, is_finished) VALUES (?, ?);
