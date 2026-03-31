-- name: CreateChapter :one
INSERT INTO chapters (title, body, user_id, status)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetChaptersByUser :many
SELECT * FROM chapters
WHERE user_id = $1;

-- name: GetChaptersById :one
SELECT * FROM chapters
WHERE id = $1;

-- name: GetVocabularyOfChapter :many
SELECT * FROM vocabulary
WHERE chapter_id = $1;

-- name: UpdateChapters :one
UPDATE chapters SET
title = $2 , body = $3 , status = $4
WHERE id = $1
RETURNING *;

-- name: DeleteChapter :exec
DELETE FROM chapters 
WHERE  id = $1;