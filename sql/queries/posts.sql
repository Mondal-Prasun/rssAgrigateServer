-- name: CreatePost :one
INSERT INTO posts (id,createdAt,updatedAt,title,description,publishedAt,url,feedId)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

