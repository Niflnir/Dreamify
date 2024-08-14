-- +goose Up
ALTER TABLE posts ADD COLUMN image_url VARCHAR(1024) DEFAULT '';

-- +goose Down
ALTER TABLE posts DROP COLUMN image_url;

