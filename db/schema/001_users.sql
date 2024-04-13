-- +goose up
CREATE TABLE IF NOT EXISTS users (
  id          UUID PRIMARY KEY,
  name        TEXT NOT NULL,
  created_at  TIMESTAMP NOT NULL,
  updated_at  TIMESTAMP NOT NULL
);

-- +goose down
DROP TABLE users;
