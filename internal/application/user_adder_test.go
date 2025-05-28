package application

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"microblogging/internal/domain"
	"microblogging/internal/domain/ports/mocks"
	"testing"
)

func TestUserAdderImpl_AddUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepositoryMock := mocks.NewMockUserRepository(ctrl)
	userAdder := NewUserAdder(userRepositoryMock)
	ctx := context.Background()

	t.Run("successful user addition", func(t *testing.T) {
		userID := "user123"
		expectedUser := domain.User{
			UserID: userID,
		}

		userRepositoryMock.EXPECT().
			Save(ctx, expectedUser).
			Return(nil)

		err := userAdder.AddUser(ctx, userID)

		assert.NoError(t, err)
	})

	t.Run("repository error", func(t *testing.T) {
		userID := "user123"
		expectedUser := domain.User{
			UserID: userID,
		}
		repoError := errors.New("repository error")

		userRepositoryMock.EXPECT().
			Save(ctx, expectedUser).
			Return(repoError)

		err := userAdder.AddUser(ctx, userID)

		assert.Error(t, err)
		assert.Equal(t, repoError, err)
	})
}
