package sqlc

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestSaveUser(t *testing.T) {
	arg := SaveUserParams{
		Name: sql.NullString{
			Valid:  true,
			String: "Jack",
		},
		Login: "jake_buffalo",
		Email: sql.NullString{
			Valid:  true,
			String: "jack@email.com",
		},
		PasswordHash: "password_hash",
		Phone: sql.NullString{
			Valid:  true,
			String: "77058113795",
		},
		Roles: "user,admin",
	}

	user, err := testQueries.SaveUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Login, user.Login)
	require.Equal(t, arg.Phone, user.Phone)
	require.Equal(t, arg.Roles, user.Roles)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.PasswordHash, user.PasswordHash)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.UpdatedAt)
	require.NotZero(t, user.UpdatedAt)
}

func TestUserGet(t *testing.T) {
	id, err := uuid.Parse("16aba4b7-c928-4bc9-b80a-5afca8205ca5")
	if err != nil {
		log.Fatalln("can not parse id:", err)
	}
	user, err := testQueries.GetUser(context.Background(), id)
	if err != nil {
		log.Fatalln("can not get user by id:", err)
	}
	require.Equal(t, user.ID, id)
}
