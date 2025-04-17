-- name: GetTodoByUid: many
SELECT *
FROM todos
WHERE user_id = $1;