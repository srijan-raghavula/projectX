-- name: CreateUser :one
INSERT INTO users(firstName, lastName, username, email, password, region, about, pfpURL, gender, isPrivate, site, dob)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING *;

-- name: UpdateUserProfile :one
UPDATE users
SET
    firstName = $2,
    lastName = $3,
    username = $4,
    email = $5,
    password = $6,
    region = $7,
    about = $8,
    pfpURL = $9,
    updatedAt = CURRENT_TIMESTAMP,
    gender = $10,
    isPrivate = $11,
    site = $12,
    isDeleted = $13,
    dob = $14
WHERE id = $1
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1 AND isDeleted = FALSE;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username= $1 AND isDeleted = FALSE;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 AND isDeleted = FALSE;

-- name: DelteUser :exec
UPDATE users
SET isDeleted = TRUE WHERE id = $1;

-- name: SearchUsers :many
SELECT * FROM users
WHERE (LOWER(username) LIKE LOWER('%' || $1 || '%')
   OR LOWER(firstName) LIKE LOWER('%' || $1 || '%')
   OR LOWER(lastName) LIKE LOWER('%' || $1 || '%'))
AND isDeleted = FALSE
LIMIT 10;
