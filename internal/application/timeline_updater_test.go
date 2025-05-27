package application

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"microblogging/internal/domain"
	"microblogging/internal/domain/ports/mocks"
	"testing"
	"time"
)

func TestTimelineUpdater_Update(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type input struct {
		ctx    context.Context
		userID string
		tweet  domain.Tweet
	}

	timelineRepositoryMock := mocks.NewMockTimelineRepository(ctrl)
	u := NewTimelineUpdater(timelineRepositoryMock)

	createdAt := time.Now()
	tests := []struct {
		name  string
		input input
		mock  func()
		err   error
	}{
		{
			name: "Should return nil when update a timeline",
			input: input{
				ctx:    ctx,
				userID: "user-uuid",
				tweet: domain.Tweet{
					TweetID:   "tweet-uuid",
					UserID:    "user-uuid",
					Content:   "Lorem ipsum",
					CreatedAt: createdAt,
				},
			},
			mock: func() {
				timeline := domain.Timeline{
					UserID: "user-uuid",
					Limit:  10,
					Total:  1,
					Tweets: []domain.Tweet{},
				}

				timelineRepositoryMock.EXPECT().Get(ctx, "user-uuid").Return(timeline, nil)
				timeline.AddTweet(domain.Tweet{
					TweetID:   "tweet-uuid",
					UserID:    "user-uuid",
					Content:   "Lorem ipsum",
					CreatedAt: createdAt,
				})
				timelineRepositoryMock.EXPECT().Update(ctx, timeline).Return(nil)
			},
			err: nil,
		},
		{
			name: "Should return an error when update a timeline",
			input: input{
				ctx:    ctx,
				userID: "user-uuid",
				tweet: domain.Tweet{
					TweetID:   "tweet-uuid",
					UserID:    "user-uuid",
					Content:   "Lorem ipsum",
					CreatedAt: createdAt,
				},
			},
			mock: func() {
				timeline := domain.Timeline{
					UserID: "user-uuid",
					Limit:  10,
					Total:  1,
					Tweets: []domain.Tweet{},
				}

				timelineRepositoryMock.EXPECT().Get(ctx, "user-uuid").Return(timeline, nil)
				timeline.AddTweet(domain.Tweet{
					TweetID:   "tweet-uuid",
					UserID:    "user-uuid",
					Content:   "Lorem ipsum",
					CreatedAt: createdAt,
				})
				timelineRepositoryMock.EXPECT().Update(ctx, timeline).Return(errors.New("internal server error"))
			},
			err: errors.New("internal server error"),
		},
		{
			name: "Should return an error when get a timeline",
			input: input{
				ctx:    ctx,
				userID: "user-uuid",
				tweet: domain.Tweet{
					TweetID:   "tweet-uuid",
					UserID:    "user-uuid",
					Content:   "Lorem ipsum",
					CreatedAt: createdAt,
				},
			},
			mock: func() {
				timelineRepositoryMock.EXPECT().Get(ctx, "user-uuid").Return(domain.Timeline{}, errors.New("internal server error"))
			},
			err: errors.New("internal server error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := u.Update(tt.input.ctx, tt.input.userID, tt.input.tweet)
			assert.Equal(t, tt.err, err)
		})
	}
}
