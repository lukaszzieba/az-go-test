-- name: CreateUser :one
INSERT INTO users (email)
VALUES (
    $1
)
RETURNING *;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: DeleteAllUsers :exec
DELETE FROM users;
