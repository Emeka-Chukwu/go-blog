-- name: GetComments :many
SELECT * FROM comments WHERE post_id = $1;


-- name: GetCommentById :one
SELECT * FROM comments WHERE id = $1;



-- name: UpdateComment :one
UPDATE comments
SET
 content= COALESCE(sqlc.narg(content), content)
WHERE
  id = sqlc.arg(id) AND user_id = $1
RETURNING *;



-- name: CreateComment :one
INSERT INTO comments (
    id,
  post_id,
  user_id,
  content
) VALUES (
  $1, $2, $3, $4
) RETURNING *;


-- name: DeleteComment :exec
DELETE FROM comments WHERE id = $1;