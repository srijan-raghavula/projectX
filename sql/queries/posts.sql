-- name: CreatePost :one
INSERT INTO posts (author, title, body, media, isPrivate)
VALUES ($1, $2, $3, $4, $5)
    RETURNING *;

-- name: RemovePost :exec
UPDATE posts
SET isDeleted = TRUE
WHERE id = $1;

-- name: UpdatePost :one
UPDATE posts
SET
    title = $2,
    body = $3,
    media = $4,
    isPrivate = $5
WHERE id = $1
RETURNING *;

-- name: FlagPost :exec
UPDATE posts
SET
    flagged = TRUE,
    flag = $2,
    isDeleted = TRUE
WHERE id = $1;

-- name: GetPostByID :one
SELECT * FROM posts WHERE id = $1 AND isDeleted = FALSE;

-- name: GetPostByUser :many
SELECT * FROM posts
WHERE author = $1
ORDER BY createdAt DESC;

