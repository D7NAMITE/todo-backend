-- name: GetTodoByClerkID :many
SELECT todos.*
FROM todos
WHERE todos.clerk_id = $1;

-- name: CreateTodoByClerkID :exec
INSERT INTO "todos" (clerk_id, title, description, status, created_at, updated_at)
VALUES ($1, $2, $3, $4, NOW(), NOW());

-- name: DeleteTodoByTodoID :exec
DELETE FROM "todos"
WHERE todo_id = $1;

-- name: UpdateTodoByTodoID :exec
UPDATE "todos"
SET
    title = COALESCE($1, title),
    description = COALESCE($2, description),
    status = COALESCE($3, status),
    updated_at = NOW()
WHERE
    todo_id = $4;
