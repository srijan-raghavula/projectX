-- name: AddFollows :exec
INSERT INTO follows (follower_id, following_id)
VALUES ($1, $2);

-- name: RemoveFollows :exec
DELETE FROM follows
WHERE follower_id = $1 AND following_id = $2;

-- name: GetFollowingList :many
SELECT following_id FROM follows
WHERE follower_id = $1;

-- name: GetFollowerList :many
SELECT follower_id FROM follows
WHERE following_id = $1;

-- name: GetFollowerCount :one
SELECT COUNT(*) FROM follows
WHERE following_id = $1;

-- name: GetFollowingCount :one
SELECT COUNT(*) FROM follows
WHERE follower_id = $1;
