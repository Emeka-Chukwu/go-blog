// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error)
	CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error)
	CreatePost(ctx context.Context, arg CreatePostParams) (Post, error)
	CreateTags(ctx context.Context, name string) (Tag, error)
	CreateTagsToPost(ctx context.Context, arg CreateTagsToPostParams) (PostTag, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteCategory(ctx context.Context, id int32) error
	DeleteComment(ctx context.Context, id int32) error
	DeletePosts(ctx context.Context, id int32) error
	DeleteTag(ctx context.Context, id int32) error
	DissociatePostZFromTag(ctx context.Context, arg DissociatePostZFromTagParams) error
	GetCategories(ctx context.Context) ([]Category, error)
	GetCategoryById(ctx context.Context, id int32) (Category, error)
	GetCommentById(ctx context.Context, id int32) (Comment, error)
	GetComments(ctx context.Context, postID int32) ([]Comment, error)
	GetPostById(ctx context.Context, id int32) (Post, error)
	GetPosts(ctx context.Context, arg GetPostsParams) ([]Post, error)
	GetTagId(ctx context.Context, id int32) (Tag, error)
	GetTags(ctx context.Context) ([]Tag, error)
	GetUserByEmail(ctx context.Context, email sql.NullString) (User, error)
	GetUserById(ctx context.Context, id int32) (User, error)
	GetUsers(ctx context.Context) ([]User, error)
	ListPostWithComment(ctx context.Context, arg ListPostWithCommentParams) ([]ListPostWithCommentRow, error)
	ListPostWithCommentAndTags(ctx context.Context) ([]ListPostWithCommentAndTagsRow, error)
	ListPostbyCategories(ctx context.Context, id int32) ([]ListPostbyCategoriesRow, error)
	ListPostbyTag(ctx context.Context, tagID sql.NullInt32) ([]Post, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error)
	UpdateComment(ctx context.Context, arg UpdateCommentParams) (Comment, error)
	UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error)
	UpdateTag(ctx context.Context, arg UpdateTagParams) (Tag, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)