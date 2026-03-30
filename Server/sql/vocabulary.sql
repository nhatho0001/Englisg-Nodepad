-- name: GetCharacterVocabulary :many
SELECT * FROM vocabulary
WHERE  chapter_id = $1 
ORDER BY created_at  DESC
LIMIT $2 OFFSET $3;

-- name: GetVocabularyOfUser :many
SELECT c.* ,  v.id ,  v.origin_content , v.description , v.practice_time , v.created_at FROM chapters as c
INNER JOIN vocabulary as v 
ON c.id = v.chapter_id
WHERE  c.user_id = $1 
ORDER BY v.created_at  DESC
LIMIT $2 OFFSET $3;


-- name: CreateVocabulary :one
INSERT INTO vocabulary (chapter_id , origin_content , description)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateVocabulary :one
UPDATE vocabulary SET 
origin_content = $2 , description = $3
WHERE id = $1
RETURNING *;