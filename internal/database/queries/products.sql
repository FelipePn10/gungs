-- name: CreateProduct :one
INSERT INTO products (users_id, metadata, description, price, tags, location)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetProductByID :one
SELECT * FROM products
WHERE id = $1;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY created_at DESC;

-- name: ListProductsByUser :many
SELECT * FROM products
WHERE users_id = $1
ORDER BY created_at DESC;

-- name: ListProductsByTag :many
SELECT * FROM products
WHERE tags @> ARRAY[$1]::text[]
ORDER BY created_at DESC;

-- name: ListActiveProducts :many
SELECT * FROM products
WHERE created_at > NOW() - INTERVAL '30 days'
ORDER BY created_at DESC;

-- name: ListProductsByUserAndLocation :many
SELECT * FROM products
WHERE users_id = $1 AND location = $2
ORDER BY created_at DESC;

-- name: UpdateProduct :one
UPDATE products
SET 
    metadata = $1,
    description = $2,
    price = $3,
    tags = $4,
    location = $5
WHERE id = $6
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;
