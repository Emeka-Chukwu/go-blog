-- name: GetCategories :many
SELECT * FROM categories;


-- name: GetCategoryById :one
SELECT * FROM categories WHERE id = $1;



-- name: UpdateCategory :one
UPDATE categories
SET
 name = $1
WHERE
  id = $2

RETURNING *;



-- name: CreateCategory :one
INSERT INTO categories (
  id,
  name
) VALUES (
  $1, $2
) RETURNING *;


-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = $1;