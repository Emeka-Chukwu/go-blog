package db

import (
	"blog-api/util"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomComment(t *testing.T) Comment {
	post, _, _ := createRandomPost(t)
	arg := CreateCommentParams{
		PostID:  post.ID,
		UserID:  post.AuthorID,
		Content: util.RandomString(30),
		ID:      int32(util.RandomInt(1, 1000000000)),
	}
	comment, err := testQueries.CreateComment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, comment)
	require.Equal(t, arg.Content, comment.Content)
	require.NotZero(t, arg.ID, comment.ID)
	require.Equal(t, arg.PostID, comment.PostID)
	require.Equal(t, arg.UserID, comment.UserID)
	require.NotZero(t, arg.UserID, comment.UserID, comment.PostID, comment.ID)
	require.NotZero(t, comment.CreatedAt)
	return comment
}

func createRandomComment2(t *testing.T, post Post) {
	arg := CreateCommentParams{
		PostID:  post.ID,
		UserID:  post.AuthorID,
		Content: util.RandomString(30),
		ID:      int32(util.RandomInt(1, 1000000000)),
	}
	comment, err := testQueries.CreateComment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, comment)
	require.Equal(t, arg.Content, comment.Content)
	require.NotZero(t, arg.ID, comment.ID)
	require.Equal(t, arg.PostID, comment.PostID)
	require.Equal(t, arg.UserID, comment.UserID)
	require.NotZero(t, arg.UserID, comment.UserID, comment.PostID, comment.ID)
	require.NotZero(t, comment.CreatedAt)

}

func TestCreateComment(t *testing.T) {
	createRandomComment(t)
}

func TestGetCommentById(t *testing.T) {
	comment1 := createRandomComment(t)
	comment2, err := testQueries.GetCommentById(context.Background(), comment1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, comment2)
	require.Equal(t, comment1.Content, comment2.Content)
	require.Equal(t, comment1.UserID, comment2.UserID)
	require.NotZero(t, comment1.ID, comment1.UserID, comment1.PostID)
	require.NotZero(t, comment2.ID, comment2.UserID, comment2.PostID)
	require.WithinDuration(t, comment1.CreatedAt, comment1.CreatedAt, time.Second)
}

func TestUpdateComment(t *testing.T) {
	oldComment := createRandomComment(t)
	newContent := util.RandomString(8)

	updatedComment, err := testQueries.UpdateComment(context.Background(), UpdateCommentParams{
		Content: sql.NullString{Valid: true, String: newContent},
		ID:      oldComment.ID,
		UserID:  oldComment.UserID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedComment)
	require.NotEqual(t, oldComment.Content, updatedComment.Content)
	require.Equal(t, oldComment.ID, updatedComment.ID)
	require.NotZero(t, updatedComment.ID, oldComment.ID)
	require.WithinDuration(t, oldComment.CreatedAt, updatedComment.CreatedAt, time.Second*2)

}

func TestDeleteComent(t *testing.T) {
	comment1 := createRandomComment(t)
	err := testQueries.DeleteComment(context.Background(), comment1.ID)
	require.NoError(t, err)
	comment2, err := testQueries.GetCommentById(context.Background(), comment1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, comment2)
}

func TestListComments(t *testing.T) {
	post, _, _ := createRandomPost(t)
	for i := 0; i < 10; i++ {
		createRandomComment2(t, post)
	}
	comments, err := testQueries.GetComments(context.TODO(), post.ID)
	require.NoError(t, err)
	require.NotEmpty(t, comments)
	for _, comment := range comments {
		require.NotEmpty(t, comment)
		require.Equal(t, comment.PostID, post.ID)

	}
}
