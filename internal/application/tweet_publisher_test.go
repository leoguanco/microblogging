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
	p := NewTweetCreator(tweetRepositoryMock, eventBusMock)

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
		{
			name: "Should return an error if content is too long",
			input: input{
				ctx:       ctx,
				tweetID:   "tweet-uuid",
				userID:    "user-uuid",
				content:   "But I must explain to you how all this mistaken idea of denouncing pleasure and praising pain was born and I will give you a complete account of the system, and expound the actual teachings of the great explorer of the truth, the master-builder of human happiness. No one rejects, dislikes, or avoids pleasure itself, because it is pleasure, but because those who do not know how to pursue pleasure rationally encounter consequences that are extremely painful. Nor again is there anyone who loves or pursues or desires to obtain pain of itself, because it is pain, but because occasionally circumstances occur in which toil and pain can procure him some great pleasure. To take a trivial example, which of us ever undertakes laborious physical exercise, except to obtain some advantage from it? But who has any right to find fault with a man who chooses to enjoy a pleasure that has no annoying consequences, or one who avoids a pain that produces no resultant pleasure?",
				createdAt: createdAt,
			},
			mock: func() {},
			err:  errors.New("content is too long"),
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
