-- name: GetAllTodos :many
SELECT id, message, is_finished FROM todos
ORDER BY id;
