package db

import (
	"blog-api/util"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	// "github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomPassword())
	require.NoError(t, err)
	arg := CreateUserParams{
		Username: sql.NullString{
			Valid:  true,
			String: util.RandomUsername(),
		},
		Email: sql.NullString{
			Valid:  true,
			String: util.RandomEmail(),
		},
		Password: sql.NullString{
			Valid:  true,
			String: hashedPassword,
		},
		Role: sql.NullString{
			Valid: true, String: "user",
		},
		ID: int32(util.RandomInt(1, 1000000000)),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username.String, user.Username.String)
	require.Equal(t, arg.Password.String, user.Password.String)
	require.Equal(t, arg.Role.String, user.Role.String)
	require.NotZero(t, arg.Email, user.Email)

	require.NotZero(t, user.CreatedAt)
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)

}

func TestGetUserById(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUserById(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password.String, user2.Password.String)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Role.String, user2.Role.String)

}

func TestGetUserByEmail(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUserByEmail(context.Background(), user1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password.String, user2.Password.String)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Role.String, user2.Role.String)

}

func TestUpdateUserOnlyUsername(t *testing.T) {
	oldUser := createRandomUser(t)
	newUsername := util.RandomUsername()

	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		Username: sql.NullString{Valid: true, String: newUsername},
		ID:       oldUser.ID,
	})
	require.NoError(t, err)
	require.NotEqual(t, oldUser.Username.String, updatedUser.Username.String)
	require.Equal(t, oldUser.Email.String, updatedUser.Email.String)
	require.Equal(t, updatedUser.Username.String, newUsername)
	require.Equal(t, oldUser.Password.String, updatedUser.Password.String)
	require.Equal(t, oldUser.Role.String, updatedUser.Role.String)
}

func TestUpdateUserOnlyPassword(t *testing.T) {
	oldUser := createRandomUser(t)
	newPassword := util.RandomString(6)
	newHashedPassword, err := util.HashPassword(newPassword)
	require.NoError(t, err)

	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		Password: sql.NullString{
			String: newHashedPassword,
			Valid:  true,
		},
		ID: oldUser.ID,
	})
	require.NoError(t, err)
	require.NotEqual(t, oldUser.Password.String, updatedUser.Password.String)
	require.Equal(t, newHashedPassword, updatedUser.Password.String)
	require.Equal(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, oldUser.Role.String, updatedUser.Role.String)

}

func TestUpdateUserAllFields(t *testing.T) {
	oldUser := createRandomUser(t)
	newUsername := util.RandomUsername()
	newPassword := util.RandomString(6)
	newRole := util.RandomString(6)
	newHashedPassword, err := util.HashPassword(newPassword)
	require.NoError(t, err)

	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		Username: sql.NullString{Valid: true, String: newUsername},
		Password: sql.NullString{Valid: true, String: newHashedPassword},
		Role:     sql.NullString{Valid: true, String: newRole},
		ID:       oldUser.ID,
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.Password.String, updatedUser.Password.String)
	require.Equal(t, newHashedPassword, updatedUser.Password.String)
	require.Equal(t, oldUser.Email, updatedUser.Email)
	require.NotEqual(t, oldUser.Role, updatedUser.Role)
	require.Equal(t, newRole, updatedUser.Role.String)
	require.NotEqual(t, oldUser.Username, updatedUser.Username)
	require.Equal(t, newUsername, updatedUser.Username.String)

}
