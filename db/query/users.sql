-- name: GetUsers :many
SELECT * FROM users;



-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1;


-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

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



-- name: UpdateUser :one
UPDATE users
SET
 username = COALESCE(sqlc.narg(username), username),
 password= COALESCE(sqlc.narg(password), password),
 role= COALESCE(sqlc.narg(role), role)
WHERE
  id = sqlc.arg(id)
RETURNING *;