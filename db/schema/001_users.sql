-- +goose up
CREATE TABLE IF NOT EXISTS users (
  id          INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
  uuid        TEXT NOT NULL,
  name        TEXT NOT NULL,
  created_at  DATETIME NOT NULL,
  updated_at  DATETIME NOT NULL
);

-- +goose down
DROP TABLE users;
