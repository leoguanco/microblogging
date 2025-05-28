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

func TestUserFollower_Follow(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type input struct {
		ctx        context.Context
		userID     string
		followerID string
	}

	userRepositoryMock := mocks.NewMockUserRepository(ctrl)
	followeeRepositoryMock := mocks.NewMockFolloweeRepository(ctrl)
	eventBusMock := mocks.NewMockEventBus(ctrl)
	f := NewUserFollower(userRepositoryMock, followeeRepositoryMock, eventBusMock)

	tests := []struct {
		name  string
		input input
		mock  func()
		err   error
	}{
		{
			name: "Should return nil when follow a user",
			input: input{
				ctx:        ctx,
				userID:     "user-uuid",
				followerID: "follower-uuid",
			},
			mock: func() {
				user := domain.User{
					UserID:  "follower-uuid",
					Follows: []string{},
				}

				followee := domain.User{
					UserID:  "user-uuid",
					Follows: []string{""},
				}

				userRepositoryMock.EXPECT().Get(ctx, "follower-uuid").Return(user, nil)

				_ = user.Follow("user-uuid")
				userRepositoryMock.EXPECT().Save(ctx, user).Return(nil)
				userRepositoryMock.EXPECT().Get(ctx, "user-uuid").Return(followee, nil)

				followee.AddFollower("follower-uuid")
				followeeRepositoryMock.EXPECT().Save(ctx, "user-uuid", followee.GetFollowers()).Return(nil)
				userRepositoryMock.EXPECT().Save(ctx, followee).Return(nil)

				eventBusMock.EXPECT().Publish(ctx, []domain.Event{
					domain.UserFollowedDomainEvent{
						UserID:     "follower-uuid",
						FolloweeID: "user-uuid",
					},
				}).Return(nil)
			},
			err: nil,
		},
		{
			name: "Should return an error when get an user",
			input: input{
				ctx:        ctx,
				userID:     "user-uuid",
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
				userID:     "user-uuid",
				followerID: "follower-uuid",
			},
			mock: func() {
				user := domain.User{
					UserID:  "follower-uuid",
					Follows: []string{},
				}
				userRepositoryMock.EXPECT().Get(ctx, "follower-uuid").Return(user, nil)

				_ = user.Follow("user-uuid")
				userRepositoryMock.EXPECT().Save(ctx, user).Return(errors.New("internal server error"))
			},
			err: errors.New("internal server error"),
		},
		{
			name: "Should return nil when follow a user that already follows",
			input: input{
				ctx:        ctx,
				userID:     "user-uuid",
				followerID: "follower-uuid",
			},
			mock: func() {
				userRepositoryMock.EXPECT().Get(ctx, "follower-uuid").Return(domain.User{
					UserID:  "follower-uuid",
					Follows: []string{"user-uuid"},
				}, nil)
			},
			err: errors.New("already follow"),
		},
		{
			name: "Should return an error if followerID is the same as my followeeID",
			input: input{
				ctx:        ctx,
				userID:     "user-uuid",
				followerID: "user-uuid",
			},
			mock: func() {
				userRepositoryMock.EXPECT().Get(ctx, "user-uuid").Return(domain.User{
					UserID:  "user-uuid",
					Follows: []string{"following-uuid"},
				}, nil)
			},
			err: errors.New("cannot follow yourself"),
		},
		{
			name: "Should return an error when get a followee",
			input: input{
				ctx:        ctx,
				userID:     "user-uuid",
				followerID: "follower-uuid",
			},
			mock: func() {
				user := domain.User{
					UserID:  "follower-uuid",
					Follows: []string{},
				}

				userRepositoryMock.EXPECT().Get(ctx, "follower-uuid").Return(user, nil)

				_ = user.Follow("user-uuid")
				userRepositoryMock.EXPECT().Save(ctx, user).Return(nil)
				userRepositoryMock.EXPECT().Get(ctx, "user-uuid").Return(domain.User{}, errors.New("internal server error"))
			},
			err: errors.New("internal server error"),
		},
		{
			name: "Should return error when save a followee",
			input: input{
				ctx:        ctx,
				userID:     "user-uuid",
				followerID: "follower-uuid",
			},
			mock: func() {
				user := domain.User{
					UserID:  "follower-uuid",
					Follows: []string{},
				}

				followee := domain.User{
					UserID:  "user-uuid",
					Follows: []string{""},
				}

				userRepositoryMock.EXPECT().Get(ctx, "follower-uuid").Return(user, nil)

				_ = user.Follow("user-uuid")
				userRepositoryMock.EXPECT().Save(ctx, user).Return(nil)
				userRepositoryMock.EXPECT().Get(ctx, "user-uuid").Return(followee, nil)

				followee.AddFollower("follower-uuid")
				followeeRepositoryMock.EXPECT().Save(ctx, "user-uuid", followee.GetFollowers()).Return(errors.New("internal server error"))
			},
			err: errors.New("internal server error"),
		},
		{
			name: "Should return error when save user followers",
			input: input{
				ctx:        ctx,
				userID:     "user-uuid",
				followerID: "follower-uuid",
			},
			mock: func() {
				user := domain.User{
					UserID:  "follower-uuid",
					Follows: []string{},
				}

				followee := domain.User{
					UserID:  "user-uuid",
					Follows: []string{""},
				}

				userRepositoryMock.EXPECT().Get(ctx, "follower-uuid").Return(user, nil)

				_ = user.Follow("user-uuid")
				userRepositoryMock.EXPECT().Save(ctx, user).Return(nil)
				userRepositoryMock.EXPECT().Get(ctx, "user-uuid").Return(followee, nil)

				followee.AddFollower("follower-uuid")
				followeeRepositoryMock.EXPECT().Save(ctx, "user-uuid", followee.GetFollowers()).Return(nil)
				userRepositoryMock.EXPECT().Save(ctx, followee).Return(errors.New("internal server error"))
			},
			err: errors.New("internal server error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := f.Follow(tt.input.ctx, tt.input.userID, tt.input.followerID)
			assert.Equal(t, tt.err, err)
		})
	}
}
