
-- name: CreatePost :one
INSERT INTO posts (title, content)
VALUES ($1, $2)
RETURNING id, title, content, created_at;

-- name: GetPost :one
SELECT id, title, content, created_at FROM posts
WHERE id = $1;

-- name: ListPosts :many
SELECT id, title, content, created_at FROM posts
ORDER BY created_at DESC;

-- name: UpdatePost :one
UPDATE posts
SET title = $2, content = $3
WHERE id = $1
RETURNING id, title, content, created_at;




-- name: DeletePost :exec
DELETE FROM posts WHERE id = $1;
