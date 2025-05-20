-- name: CreateTodo :one
INSERT INTO todo(title, description, dueDate, status, createdAt, updatedAt) 
VALUES ($1, $2, $3, $4,now(),now())
RETURNING *;

-- name: GetAllTodo :many
SELECT * FROM todo where status = 'PENDING';

-- name: GetTodoById :one
SELECT * FROM todo WHERE id = $1;

-- name: UpdateTodo :one
UPDATE todo SET title = $1, description = $2, dueDate = $3,updatedAt = now()
WHERE id = $4
RETURNING *;

-- name: DeleteTodo :one
DELETE FROM todo WHERE id = $1
RETURNING *;

-- name: MarkTodoAsCompleted :one
UPDATE todo SET status = 'COMPLETED', updatedAt = now()
WHERE id = $1
RETURNING *;