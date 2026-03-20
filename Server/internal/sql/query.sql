-- name: GetAuthor :one
SELECT * FROM users
WHERE email = $1 and deleted_at IS NULL 
LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM users
WHERE deleted_at IS NOT NULL;

-- name: CreateUser :one
INSERT INTO users (email, hashed_password)
VALUES ($1 , $2)
RETURNING *;