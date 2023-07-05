// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: post_tags.sql

package db

import (
	"context"
	"database/sql"
)

const createTagsToPost = `-- name: CreateTagsToPost :one
INSERT INTO post_tags (
 post_id, tag_id
) VALUES (
  $1, $2
) RETURNING id, post_id, tag_id, created_at, updated_at
`

type CreateTagsToPostParams struct {
	PostID sql.NullInt32 `json:"post_id"`
	TagID  sql.NullInt32 `json:"tag_id"`
}

func (q *Queries) CreateTagsToPost(ctx context.Context, arg CreateTagsToPostParams) (PostTag, error) {
	row := q.db.QueryRowContext(ctx, createTagsToPost, arg.PostID, arg.TagID)
	var i PostTag
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.TagID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteTagsOfPost = `-- name: DeleteTagsOfPost :exec
DELETE FROM post_tags WHERE post_id = $1
`

func (q *Queries) DeleteTagsOfPost(ctx context.Context, postID sql.NullInt32) error {
	_, err := q.db.ExecContext(ctx, deleteTagsOfPost, postID)
	return err
}

const dissociatePostZFromTag = `-- name: DissociatePostZFromTag :exec
DELETE FROM post_tags WHERE post_id = $1 AND tag_id=$2
`

type DissociatePostZFromTagParams struct {
	PostID sql.NullInt32 `json:"post_id"`
	TagID  sql.NullInt32 `json:"tag_id"`
}

func (q *Queries) DissociatePostZFromTag(ctx context.Context, arg DissociatePostZFromTagParams) error {
	_, err := q.db.ExecContext(ctx, dissociatePostZFromTag, arg.PostID, arg.TagID)
	return err
}

const updateTagsPost = `-- name: UpdateTagsPost :one
UPDATE post_tags 
SET 
    tag_id = COALESCE($1,tag_id)
WHERE 
    post_id = $2

RETURNING id, post_id, tag_id, created_at, updated_at
`

type UpdateTagsPostParams struct {
	TagID  sql.NullInt32 `json:"tag_id"`
	PostID sql.NullInt32 `json:"post_id"`
}

func (q *Queries) UpdateTagsPost(ctx context.Context, arg UpdateTagsPostParams) (PostTag, error) {
	row := q.db.QueryRowContext(ctx, updateTagsPost, arg.TagID, arg.PostID)
	var i PostTag
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.TagID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
