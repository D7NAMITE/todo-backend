-- name: GetTodoByUid :many
SELECT todos.*
FROM todos
WHERE user_id = $1;

-- name: CreateTodo :exec
INSERT INTO "todos" (user_id, title, description, status, created_at, updated_at)
VALUES
    ($1, $2, $3, $4, NOW(), NOW());