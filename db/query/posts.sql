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
    id,
  title,
  content,
  author_id,
  category_id
) VALUES (
  $1, $2, $3, $4 , $5
) RETURNING *;


-- name: DeletePosts :exec
DELETE FROM posts WHERE id = $1;





-- name: ListPostWithCommentAndTags :many
SELECT p.*, json_agg(c.*) AS comments, json_agg(t.*) AS tags
FROM posts p
LEFT JOIN comments c ON c.post_id = p.id
LEFT JOIN post_tags pt ON pt.post_id = p.id
LEFT JOIN tags t ON t.id = pt.tag_id
GROUP BY p.id, p.title;


-- name: ListPostbyCategories :many
SELECT p.*, json_agg(c.*) AS comments, json_agg(t.*) AS tags
FROM posts p 
LEFT JOIN comments c ON c.post_id = p.id
LEFT JOIN post_tags pt ON pt.post_id = p.id
LEFT JOIN tags t ON t.id = pt.tag_id
Where p.category_id =  $1
GROUP BY p.id, p.title;

-- name: ListPostbyTag :many
SELECT p.*,json_agg(c.*) AS comments FROM posts p
JOIN post_tags pt ON pt.post_id = p.id
LEFT JOIN comments c ON c.post_id = p.id
WHERE pt.tag_id = $1
Group by p.id;

