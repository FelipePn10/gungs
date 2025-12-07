-- name: CreateUser :one
INSERT INTO users (email, name, age, biography, address, number)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at DESC;

-- name: UpdateUser :one
UPDATE users
SET 
    email = $1,
    name = $2,
    age = $3,
    biography = $4,
    address = $5,
    number = $6
WHERE id = $7
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
