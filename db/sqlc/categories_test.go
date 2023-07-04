package db

import (
	"blog-api/util"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) Category {
	arg := CreateCategoryParams{
		Name: sql.NullString{Valid: true, String: util.RandomString(8)},
		ID:   int32(util.RandomInt(1, 1000000000)),
	}
	category, err := testQueries.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)
	require.Equal(t, arg.Name, category.Name)
	require.NotZero(t, arg.ID, category.ID)
	require.NotZero(t, category.CreatedAt)
	return category
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategoryById(t *testing.T) {
	category1 := createRandomCategory(t)
	category2, err := testQueries.GetCategoryById(context.Background(), category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)
	require.Equal(t, category1.Name, category2.Name)
	require.NotZero(t, category1.ID)
	require.NotZero(t, category2.ID)
	require.WithinDuration(t, category1.CreatedAt.Time, category2.CreatedAt.Time, time.Second)
}

func TestUpdateCategory(t *testing.T) {
	oldCategory := createRandomCategory(t)
	newUsername := util.RandomString(8)

	updatedCategory, err := testQueries.UpdateCategory(context.Background(), UpdateCategoryParams{
		Name: sql.NullString{Valid: true, String: newUsername},
		ID:   oldCategory.ID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedCategory)
	require.NotEqual(t, oldCategory.Name, updatedCategory.Name)
	require.Equal(t, oldCategory.ID, updatedCategory.ID)
	require.NotZero(t, updatedCategory.ID)
	require.WithinDuration(t, oldCategory.CreatedAt.Time, updatedCategory.CreatedAt.Time, time.Second*2)

}

func TestDeleteCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	err := testQueries.DeleteCategory(context.Background(), category1.ID)
	require.NoError(t, err)
	category2, err := testQueries.GetCategoryById(context.Background(), category1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, category2)
}

func TestListCategories(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCategory(t)
	}
	accounts, err := testQueries.GetCategories(context.TODO())
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	for _, account := range accounts {
		require.NotEmpty(t, account)

	}
}
