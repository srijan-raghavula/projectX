-- +goose Up
CREATE TABLE likes (
    post_id UUID REFERENCES posts(id) ON DELETE CASCADE,
    user_id UUID REFERENCES posts(id) ON DELETE CASCADE,
    PRIMARY KEY (post_id, user_id)
);

-- +goose Down
DROP TABLE likes;
