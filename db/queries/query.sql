-- name: CreateUser :one
INSERT INTO users (id, uuid, name, created_at, updated_at)
VALUES (?, ?, ?, ?, ?)
RETURNING *;
