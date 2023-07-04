-- name: GetTags :many
SELECT * FROM tags;


-- name: GetTagId :one
SELECT * FROM tags WHERE id = $1;



-- name: UpdateTag :one
UPDATE tags
SET
 name = $1
WHERE
  id = $2
RETURNING *;



-- name: CreateTags :one
INSERT INTO tags (
    id,
  name
) VALUES (
  $1, $2
) RETURNING *;


-- name: DeleteTag :exec
DELETE FROM tags WHERE id = $1;