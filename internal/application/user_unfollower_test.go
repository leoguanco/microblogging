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

func TestUserUnfollower_Unfollow(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type input struct {
		ctx        context.Context
		followeeID string
		followerID string
	}

	userRepositoryMock := mocks.NewMockUserRepository(ctrl)
	eventBusMock := mocks.NewMockEventBus(ctrl)
	u := NewUserUnfollower(userRepositoryMock, eventBusMock)

	tests := []struct {
		name  string
		input input
		mock  func()
		err   error
	}{
		{
			name: "Should return nil when unfollow a user",
			input: input{
				ctx:        ctx,
				followeeID: "user-uuid",
				followerID: "follower-uuid",
			},
			mock: func() {
				user := domain.User{
					UserID:  "follower-uuid",
					Follows: []string{"user-uuid"},
				}

				userRepositoryMock.EXPECT().Get(ctx, "follower-uuid").Return(user, nil)

				user.Unfollow("user-uuid")
				userRepositoryMock.EXPECT().Save(ctx, user).Return(nil)
				eventBusMock.EXPECT().Publish(ctx, []domain.Event{
					domain.UserUnfollowedDomainEvent{
						UserID:       "follower-uuid",
						UnFolloweeID: "user-uuid",
					},
				}).Return(nil)
			},
			err: nil,
		},
		{
			name: "Should return an error when get an user",
			input: input{
				ctx:        ctx,
				followeeID: "user-uuid",
				followerID: "follower-uuid",
			},
			mock: func() {
				userRepositoryMock.EXPECT().Get(ctx, "follower-uuid").Return(domain.User{}, errors.New("internal server error"))
			},
			err: errors.New("internal server error"),
		},
		{
			name: "Should return an error when save an user",
			input: input{
				ctx:        ctx,
				followeeID: "user-uuid",
				followerID: "follower-uuid",
			},
			mock: func() {
				user := domain.User{
					UserID:  "follower-uuid",
					Follows: []string{"user-uuid"},
				}

				userRepositoryMock.EXPECT().Get(ctx, "follower-uuid").Return(user, nil)

				user.Unfollow("user-uuid")
				userRepositoryMock.EXPECT().Save(ctx, user).Return(errors.New("internal server error"))
			},
			err: errors.New("internal server error"),
		},
		{
			name: "Should return an error when publish in event bus",
			input: input{
				ctx:        ctx,
				followeeID: "user-uuid",
				followerID: "follower-uuid",
			},
			mock: func() {
				user := domain.User{
					UserID:  "follower-uuid",
					Follows: []string{"user-uuid"},
				}

				userRepositoryMock.EXPECT().Get(ctx, "follower-uuid").Return(user, nil)

				user.Unfollow("user-uuid")
				userRepositoryMock.EXPECT().Save(ctx, user).Return(nil)
				eventBusMock.EXPECT().Publish(ctx, []domain.Event{
					domain.UserUnfollowedDomainEvent{
						UserID:       "follower-uuid",
						UnFolloweeID: "user-uuid",
					},
				}).Return(errors.New("internal server error"))
			},
			err: errors.New("internal server error"),
		},
		{
			name: "Should return an error unfollow yourself",
			input: input{
				ctx:        ctx,
				followeeID: "follower-uuid",
				followerID: "follower-uuid",
			},
			mock: func() {
				user := domain.User{
					UserID:  "follower-uuid",
					Follows: []string{"user-uuid"},
				}

				userRepositoryMock.EXPECT().Get(ctx, "follower-uuid").Return(user, nil)
			},
			err: errors.New("cannot unfollow yourself"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := u.Unfollow(tt.input.ctx, tt.input.followeeID, tt.input.followerID)
			assert.Equal(t, tt.err, err)
		})
	}
}
