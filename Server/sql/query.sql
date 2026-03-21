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

-- name: CreateToken :one
INSERT INTO refresh_tokens (user_id, hashed_token , created_at , expires_at)
VALUES ($1 , $2 , $3 , $4)
RETURNING *;


-- name: GetTokensByUid :many
SELECT * FROM refresh_tokens
WHERE user_id = $1; 


-- name: DeleteUserToken :exec
DELETE FROM refresh_tokens
WHERE user_id = $1;