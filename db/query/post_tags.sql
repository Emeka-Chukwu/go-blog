



-- name: UpdateTagsPost :one
UPDATE post_tags 
SET 
    tag_id = COALESCE(sqlc.narg(tag_id),tag_id),
    post_id = COALESCE(sqlc.narg(post_id),tag_id)
WHERE 
    id = $1

RETURNING *;


UPDATE posts
SET
 title = COALESCE(sqlc.narg(title), title),
 content= COALESCE(sqlc.narg(content), content),
 category_id= COALESCE(sqlc.narg(category_id), category_id)
WHERE
  id = sqlc.arg(id) AND author_id = $1
RETURNING *;




-- name: CreateTagsToPost :one
INSERT INTO post_tags (
 post_id, tag_id
) VALUES (
  $1, $2
) RETURNING *;


-- name: DissociatePostZFromTag :exec
DELETE FROM post_tags WHERE post_id = $1 AND tag_id=$2;

-- name: DeleteTagsOfPost :exec
DELETE FROM post_tags WHERE post_id = $1;