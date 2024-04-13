-- +goose up
CREATE TABLE IF NOT EXISTS users (
  id          uuid PRIMARY KEY NOT NULL,
  name        TEXT NOT NULL,
  created_at  TIMESTAMP NOT NULL,
  updated_at  TIMESTAMP NOT NULL
);

-- +goose down
DROP TABLE users;
