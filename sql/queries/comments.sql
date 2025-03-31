-- :name AddComment :one
WITH newComment AS(
    INSERT INTO comments (post_id, user_id, comment)
    VALUES ($1, $2, $3)
        RETURNING post_id
) UPDATE posts
SET comments = comments + 1
WHERE id = (SELECT post_id FROM newCommnet);

-- :name RemoveComment :one
WITH deleteComment AS(
    DELETE FROM comments
    WHERE id = $1
        RETURNING post_id
) UPDATE posts
SET comments = comments - 1
WHERE id = (SELECT post_id FROM deleteCommnet);
