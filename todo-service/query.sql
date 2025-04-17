-- name: GetTodoByUid :many
SELECT todos.*
FROM todos
WHERE user_id = $1;