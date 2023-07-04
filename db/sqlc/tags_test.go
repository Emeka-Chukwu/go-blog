package db

import (
	"blog-api/util"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomTag(t *testing.T) Tag {
	arg := CreateTagsParams{
		Name: util.RandomString(7),
		ID:   int32(util.RandomInt(1, 1000000000)),
	}
	tag, err := testQueries.CreateTags(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, tag)
	require.Equal(t, arg.Name, tag.Name)
	require.NotZero(t, arg.ID, tag.ID)
	require.NotZero(t, tag.CreatedAt)
	return tag
}

func TestCreateTags(t *testing.T) {
	createRandomTag(t)
}

func TestGetTagById(t *testing.T) {
	tag1 := createRandomTag(t)
	tag2, err := testQueries.GetTagId(context.Background(), tag1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, tag2)
	require.Equal(t, tag1.Name, tag2.Name)
	require.NotZero(t, tag1.ID)
	require.NotZero(t, tag2.ID)
	require.WithinDuration(t, tag1.CreatedAt, tag2.CreatedAt, time.Second)
}

func TestUpdateTag(t *testing.T) {
	oldTag := createRandomTag(t)
	newName := util.RandomString(7)

	updatedTag, err := testQueries.UpdateTag(context.Background(), UpdateTagParams{
		Name: newName,
		ID:   oldTag.ID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedTag)
	require.NotEqual(t, oldTag.Name, updatedTag.Name)
	require.Equal(t, oldTag.ID, updatedTag.ID)
	require.NotZero(t, updatedTag.ID)
	require.WithinDuration(t, oldTag.CreatedAt, updatedTag.CreatedAt, time.Second*2)

}

func TestDeleteTag(t *testing.T) {
	tag1 := createRandomTag(t)
	err := testQueries.DeleteTag(context.Background(), tag1.ID)
	require.NoError(t, err)
	tag2, err := testQueries.GetTagId(context.Background(), tag1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, tag2)
}

func TestListTags(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTag(t)
	}
	tags, err := testQueries.GetTags(context.TODO())
	require.NoError(t, err)
	require.NotEmpty(t, tags)
	for _, tag := range tags {
		require.NotEmpty(t, tag)

	}
}
