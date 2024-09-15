-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS posts (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  title TEXT NOT NULL,
  description TEXT,
  published_at TIMESTAMP NOT NULL,
  url TEXT NOT NULL UNIQUE,
  feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS posts;
-- +goose StatementEnd
