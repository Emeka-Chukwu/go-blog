-- name: GetPosts :many
SELECT * FROM posts ORDER BY id
LIMIT $1
OFFSET $2;


-- name: GetPostById :one
SELECT * FROM posts WHERE id = $1;



-- name: UpdatePost :one
UPDATE posts
SET
 title = COALESCE(sqlc.narg(title), title),
 content= COALESCE(sqlc.narg(content), content),
 category_id= COALESCE(sqlc.narg(category_id), category_id)
WHERE
  id = sqlc.arg(id) AND author_id = $1
RETURNING *;



-- name: CreatePost :one
INSERT INTO posts (
  title,
  content,
  author_id,
  category_id
) VALUES (
  $1, $2, $3, $4 
) RETURNING *;


-- name: DeletePosts :exec
DELETE FROM posts WHERE id = $1;



-- name: ListPostWithComment :many
SELECT p.*, json_agg(c.*) AS comments
FROM posts p
LEFT JOIN comments c ON c.post_id = p.id
GROUP BY p.id LIMIT $1
OFFSET $2;

-- name: ListPostWithCommentAndTags :many
SELECT p.*, json_agg(c.*) AS comments, json_agg(t.tags) AS tags
FROM posts p
LEFT JOIN comments c ON c.post_id = p.id
LEFT JOIN post_tags pt ON pt.post_id = p.id
LEFT JOIN tags t ON t.id = pt.tag_id
GROUP BY p.id, p.title;

