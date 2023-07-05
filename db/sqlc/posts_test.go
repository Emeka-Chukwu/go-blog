package db

import (
	"blog-api/util"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomPost(t *testing.T) (Post, Category, Tag, PostTag) {
	category := createRandomCategory(t)
	user := createRandomUser(t)
	tag := createRandomTag(t)
	arg := CreatePostParams{
		Title:      util.RandomString(24),
		ID:         int32(util.RandomInt(1, 1000000000)),
		Content:    util.RandomString(200),
		AuthorID:   user.ID,
		CategoryID: category.ID,
	}
	post, err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post)
	postTagArg := CreateTagsToPostParams{
		PostID: sql.NullInt32{Valid: true, Int32: post.ID},
		TagID:  sql.NullInt32{Valid: true, Int32: tag.ID},
	}
	postTag, err := testQueries.CreateTagsToPost(context.Background(), postTagArg)
	require.NoError(t, err)
	require.NotEmpty(t, postTag)
	require.Equal(t, arg.Title, post.Title)
	require.Equal(t, arg.Content, post.Content)
	require.Equal(t, arg.AuthorID, post.AuthorID)
	require.Equal(t, arg.CategoryID, post.CategoryID)
	require.NotZero(t, arg.ID, category.ID)
	require.NotZero(t, category.CreatedAt)
	require.Equal(t, postTagArg.PostID.Int32, post.ID)
	require.Equal(t, postTagArg.TagID.Int32, tag.ID)
	return post, category, tag, postTag
}

func TestCreatePost(t *testing.T) {
	createRandomPost(t)
}

func TestGetPostById(t *testing.T) {
	post1, _, _, _ := createRandomPost(t)
	post2, err := testQueries.GetPostById(context.Background(), post1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, post2)
	require.Equal(t, post1.Title, post2.Title)
	require.NotZero(t, post1.ID)
	require.NotZero(t, post2.ID)
	require.WithinDuration(t, post1.CreatedAt, post2.CreatedAt, time.Second)
}
func TestGetPostByIdFailed(t *testing.T) {
	post, err := testQueries.GetPostById(context.Background(), 0000)
	require.Error(t, err)
	require.Empty(t, post)
	require.Zero(t, post.ID)

}

func TestUpdatePostContentonly(t *testing.T) {
	oldPost, _, _, _ := createRandomPost(t)
	newContent := util.RandomString(200)

	updatedPost, err := testQueries.UpdatePost(context.Background(), UpdatePostParams{
		Content:  sql.NullString{Valid: true, String: newContent},
		ID:       oldPost.ID,
		AuthorID: oldPost.AuthorID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedPost)
	require.NotEqual(t, oldPost.Content, updatedPost.Content)
	require.Equal(t, oldPost.Title, updatedPost.Title)
	require.NotZero(t, updatedPost.ID)
	require.WithinDuration(t, oldPost.CreatedAt, updatedPost.CreatedAt, time.Second*2)

}

func TestUpdatePost(t *testing.T) {
	oldPost, oldCategory, oldTag, postTag := createRandomPost(t)
	newContent := util.RandomString(200)
	newTitle := util.RandomString(20)
	newCategory := createRandomCategory(t)
	newTag := createRandomTag(t)

	updatedPost, err := testQueries.UpdatePost(context.Background(), UpdatePostParams{
		Content:    sql.NullString{Valid: true, String: newContent},
		ID:         oldPost.ID,
		Title:      sql.NullString{Valid: true, String: newTitle},
		CategoryID: sql.NullInt32{Valid: true, Int32: newCategory.ID},
		AuthorID:   oldPost.AuthorID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedPost)
	updatedTag, err := testQueries.UpdateTagsPost(context.Background(), UpdateTagsPostParams{
		PostID: sql.NullInt32{Valid: true, Int32: oldPost.ID},
		TagID:  sql.NullInt32{Valid: true, Int32: newTag.ID},
		ID:     postTag.ID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedTag)

	require.NotEqual(t, oldPost.Content, updatedPost.Content)
	require.NotEqual(t, oldPost.Title, updatedPost.Title)
	require.NotEqual(t, oldPost.CategoryID, updatedPost.CategoryID)
	require.Equal(t, oldPost.ID, updatedPost.ID)
	////tag
	require.NotEqual(t, oldTag.ID, updatedTag.ID)
	require.NotZero(t, oldTag.ID, updatedTag.ID)
	require.NotZero(t, updatedTag.ID)
	require.NotZero(t, oldTag.ID)
	require.NotZero(t, updatedPost.ID)
	///category

	require.NotZero(t, newCategory.ID)
	require.NotZero(t, oldCategory.ID)

	require.WithinDuration(t, oldPost.CreatedAt, updatedPost.CreatedAt, time.Second*2)

}

func TestDeletePost(t *testing.T) {
	post1, _, _, _ := createRandomPost(t)
	err := testQueries.DeleteTagsOfPost(context.Background(), sql.NullInt32{Valid: true, Int32: post1.ID})
	require.NoError(t, err)
	err = testQueries.DeletePosts(context.Background(), post1.ID)
	require.NoError(t, err)
	category2, err := testQueries.GetPostById(context.Background(), post1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, category2)
}

func TestDisociatePostFromTag(t *testing.T) {
	post1, _, tag, postTag := createRandomPost(t)
	err := testQueries.DissociatePostZFromTag(context.Background(), DissociatePostZFromTagParams{
		PostID: sql.NullInt32{Valid: true, Int32: post1.ID},
		TagID:  sql.NullInt32{Valid: true, Int32: tag.ID},
	})
	require.NoError(t, err)
	updatedTag, err := testQueries.UpdateTagsPost(context.Background(), UpdateTagsPostParams{
		PostID: sql.NullInt32{Valid: true, Int32: post1.ID},
		TagID:  sql.NullInt32{Valid: true, Int32: tag.ID},
		ID:     postTag.ID,
	})

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, updatedTag)
}

func TestListPosts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomPost(t)
	}
	posts, err := testQueries.GetPosts(context.TODO(), GetPostsParams{Limit: 10, Offset: 1})
	require.NoError(t, err)
	require.NotEmpty(t, posts)
	for _, post := range posts {
		require.NotEmpty(t, post)

	}
}

// ListPostWithCommentAndTags

func TestListPostsWithTagsAndComments(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomPost(t)
	}
	posts, err := testQueries.ListPostWithCommentAndTags(context.TODO())
	require.NoError(t, err)
	require.NotEmpty(t, posts)

	for _, post := range posts {
		require.NotEmpty(t, post)

	}
}

func TestListPostbyCategory(t *testing.T) {
	var lastCategory Category
	for i := 0; i < 10; i++ {
		_, category, _, _ := createRandomPost(t)
		lastCategory = category
	}
	posts, err := testQueries.ListPostbyCategories(context.TODO(), lastCategory.ID)
	require.NoError(t, err)
	require.NotEmpty(t, posts)
	for _, post := range posts {
		require.NotEmpty(t, post)
		require.NotEmpty(t, post)
		require.Equal(t, lastCategory.ID, post.CategoryID)
	}
}

func TestListPostbyTag(t *testing.T) {
	var lastTag Tag
	for i := 0; i < 10; i++ {
		_, _, tag, _ := createRandomPost(t)
		lastTag = tag
	}
	posts, err := testQueries.ListPostbyTag(context.TODO(), sql.NullInt32{Valid: true, Int32: lastTag.ID})
	require.NoError(t, err)
	require.NotEmpty(t, posts)
	for _, post := range posts {
		require.NotEmpty(t, post)
		require.NotEmpty(t, post)

	}
}
