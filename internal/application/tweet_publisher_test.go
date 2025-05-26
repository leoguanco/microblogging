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

func TestTweetPublisher_Publish(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type input struct {
		ctx       context.Context
		tweetID   string
		userID    string
		content   string
		createdAt time.Time
	}

	tweetRepositoryMock := mocks.NewMockTweetRepository(ctrl)
	eventBusMock := mocks.NewMockEventBus(ctrl)
	p := NewTweetPublisher(tweetRepositoryMock, eventBusMock)

	createdAt := time.Now()
	inputTweet, _ := domain.NewTweet("tweet-uuid", "user-uuid", "tweet content", createdAt)

	tests := []struct {
		name  string
		input input
		mock  func()
		err   error
	}{
		{
			name: "Should return nil when publish a tweet",
			input: input{
				ctx:       ctx,
				tweetID:   "tweet-uuid",
				userID:    "user-uuid",
				content:   "tweet content",
				createdAt: createdAt,
			},
			mock: func() {
				tweetRepositoryMock.EXPECT().Save(ctx, *inputTweet).Return(nil)
				eventBusMock.EXPECT().Publish(ctx, []domain.Event{
					{
						"AggregateID": "tweet-uuid",
						"Type":        "tweet_created",
						"UserID":      "user-uuid",
						"Content":     "tweet content",
						"CreatedAt":   createdAt,
					},
				}).Return(nil)
			},
			err: nil,
		},
		{
			name: "Should return an error when publish a tweet",
			input: input{
				ctx:       ctx,
				tweetID:   "tweet-uuid",
				userID:    "user-uuid",
				content:   "tweet content",
				createdAt: createdAt,
			},
			mock: func() {
				tweetRepositoryMock.EXPECT().Save(ctx, *inputTweet).Return(errors.New("internal server error"))
			},
			err: errors.New("internal server error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := p.Create(
				tt.input.ctx,
				tt.input.tweetID,
				tt.input.userID,
				tt.input.content,
				tt.input.createdAt,
			)
			assert.Equal(t, tt.err, err)
		})
	}
}
