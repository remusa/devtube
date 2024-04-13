-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, uuid)
VALUES (?, ?, ?, ?, ?)
RETURNING *;
