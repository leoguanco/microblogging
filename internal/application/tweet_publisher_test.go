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
		ctx   context.Context
		tweet domain.Tweet
	}

	tweetRepositoryMock := mocks.NewMockTweetRepository(ctrl)
	p := NewTweetPublisher(tweetRepositoryMock)

	createdAt := time.Now()
	inputTweet := domain.NewTweet("tweet-uuid", "user-uuid", "tweet content", createdAt)

	tests := []struct {
		name  string
		input input
		mock  func()
		err   error
	}{
		{
			name: "Should return nil when publish a tweet",
			input: input{
				ctx:   ctx,
				tweet: inputTweet,
			},
			mock: func() {
				tweetRepositoryMock.EXPECT().Save(ctx, inputTweet).Return(nil)
			},
			err: nil,
		},
		{
			name: "Should return an error when publish a tweet",
			input: input{
				ctx:   ctx,
				tweet: inputTweet,
			},
			mock: func() {
				tweetRepositoryMock.EXPECT().Save(ctx, inputTweet).Return(errors.New("internal server error"))
			},
			err: errors.New("internal server error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := p.Publish(tt.input.ctx, tt.input.tweet)
			assert.Equal(t, tt.err, err)
		})
	}
}
