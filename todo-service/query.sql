-- name: GetTodoByClerkID :many
SELECT todos.*
FROM todos
JOIN users ON todos.user_id = users.user_id
WHERE users.clerk_id = $1;

-- name: CreateTodoByClerkID :exec
INSERT INTO "todos" (user_id, title, description, status, created_at, updated_at)
VALUES ((SELECT user_id FROM users WHERE clerk_id = $1),
        $2, $3, $4, NOW(), NOW());