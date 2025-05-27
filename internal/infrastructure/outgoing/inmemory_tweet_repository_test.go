package outgoing

import (
	"context"
	"microblogging/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryTweetRepository_Save(t *testing.T) {
	type input struct {
		initialData map[string]domain.Tweet
		tweet       domain.Tweet
	}

	type want struct {
		err bool
	}

	now := time.Now()

	tests := []struct {
		name  string
		input input
		want  want
	}{
		{
			name: "Save new tweet",
			input: input{
				initialData: map[string]domain.Tweet{},
				tweet: domain.Tweet{
					TweetID:   "tweet1",
					Content:   "Hello World",
					UserID:    "user1",
					CreatedAt: now,
				},
			},
			want: want{
				err: false,
			},
		},
		{
			name: "Save tweet with existing ID",
			input: input{
				initialData: map[string]domain.Tweet{
					"tweet1": {
						TweetID:   "tweet1",
						Content:   "Original Content",
						UserID:    "user1",
						CreatedAt: now,
					},
				},
				tweet: domain.Tweet{
					TweetID:   "tweet1",
					Content:   "Updated Content",
					UserID:    "user1",
					CreatedAt: now,
				},
			},
			want: want{
				err: false,
			},
		},
		{
			name: "Save tweet with empty ID",
			input: input{
				initialData: map[string]domain.Tweet{},
				tweet: domain.Tweet{
					TweetID:   "",
					Content:   "Hello World",
					UserID:    "user1",
					CreatedAt: now,
				},
			},
			want: want{
				err: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &InMemoryTweetRepository{
				tweets: make(map[string]domain.Tweet),
			}

			for k, v := range tt.input.initialData {
				repo.tweets[k] = v
			}

			err := repo.Save(context.Background(), tt.input.tweet)

			if tt.want.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				storedTweet, exists := repo.tweets[tt.input.tweet.TweetID]
				assert.True(t, exists)
				assert.Equal(t, tt.input.tweet, storedTweet)
			}
		})
	}
}
