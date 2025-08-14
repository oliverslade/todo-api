-- name: GetAllTodos :many
SELECT id, message, is_finished FROM todos
ORDER BY id;

-- name: CreateTodo :exec
INSERT INTO todos (message, is_finished) VALUES (?, ?);
