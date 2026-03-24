-- name: CreateChapter :one
INSERT INTO chapters (title, body, user_id, status)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetChaptersByUser :many
SELECT * FROM chapters
WHERE user_id = $1;