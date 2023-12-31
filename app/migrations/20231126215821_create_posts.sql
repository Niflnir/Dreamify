-- +goose Up
CREATE TABLE posts (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  body TEXT NOT NULL,
  date_created date NOT NULL DEFAULT CURRENT_DATE
);

-- +goose Down
DROP TABLE posts;

