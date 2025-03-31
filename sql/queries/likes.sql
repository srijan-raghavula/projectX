-- :name AddLike :exec
WITH addLike AS (
    INSERT INTO likes (post_id, user_id)
    VALUES ($1, $2)
        RETURNING post_id
) UPDATE posts
SET likes = likes + 1
WHERE id = (SELECT post_id FROM addLike);

-- :name RemoveLike :exec
WITH removeLike AS (
    DELETE FROM likes
    WHERE post_id = $1 AND user_id = $2
        RETURNING post_id
) UPDATE posts
SET likes = likes - 1
WHERE id = (SELECT post_id FROM removeLike);
