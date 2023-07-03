-- name: GetUsers :many
SELECT * FROM users;



-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1;


-- name: CreateUser :one
INSERT INTO users (
  id,
  username,
  email,
  password,
  role
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

