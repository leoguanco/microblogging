package outgoing

import (
	"context"
	"microblogging/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryTimelineRepository_Get(t *testing.T) {
	type input struct {
		initialData map[string]domain.Timeline
		userID      string
	}

	type want struct {
		timeline domain.Timeline
		err      bool
	}

	tests := []struct {
		name  string
		input input
		want  want
	}{
		{
			name: "Get existing timeline",
			input: input{
				initialData: map[string]domain.Timeline{
					"user1": {
						UserID: "user1",
						Tweets: []domain.Tweet{
							{TweetID: "tweet1", Content: "Hello World", UserID: "user1"},
						},
					},
				},
				userID: "user1",
			},
			want: want{
				timeline: domain.Timeline{
					UserID: "user1",
					Tweets: []domain.Tweet{
						{TweetID: "tweet1", Content: "Hello World", UserID: "user1"},
					},
				},
				err: false,
			},
		},
		{
			name: "Get non-existing timeline returns empty timeline",
			input: input{
				initialData: map[string]domain.Timeline{
					"user1": {
						UserID: "user1",
						Tweets: []domain.Tweet{
							{TweetID: "tweet1", Content: "Hello World", UserID: "user1"},
						},
					},
				},
				userID: "user2",
			},
			want: want{
				timeline: domain.Timeline{
					UserID: "user2",
					Tweets: []domain.Tweet{},
				},
				err: false,
			},
		},
		{
			name: "Get timeline with empty userID",
			input: input{
				initialData: map[string]domain.Timeline{},
				userID:      "",
			},
			want: want{
				timeline: domain.Timeline{
					UserID: "",
					Tweets: []domain.Tweet{},
				},
				err: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &InMemoryTimelineRepository{
				timelines: tt.input.initialData,
			}

			result, err := repo.Get(context.Background(), tt.input.userID)

			if tt.want.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.timeline, result)
			}
		})
	}
}

func TestInMemoryTimelineRepository_Update(t *testing.T) {
	type input struct {
		initialData map[string]domain.Timeline
		timeline    domain.Timeline
	}

	type want struct {
		err bool
	}

	tests := []struct {
		name  string
		input input
		want  want
	}{
		{
			name: "Update existing timeline",
			input: input{
				initialData: map[string]domain.Timeline{
					"user1": {
						UserID: "user1",
						Tweets: []domain.Tweet{
							{TweetID: "tweet1", Content: "Hello World", UserID: "user1"},
						},
					},
				},
				timeline: domain.Timeline{
					UserID: "user1",
					Tweets: []domain.Tweet{
						{TweetID: "tweet1", Content: "Hello World", UserID: "user1"},
						{TweetID: "tweet2", Content: "Hello Again", UserID: "user1"},
					},
				},
			},
			want: want{
				err: false,
			},
		},
		{
			name: "Create new timeline",
			input: input{
				initialData: map[string]domain.Timeline{},
				timeline: domain.Timeline{
					UserID: "user1",
					Tweets: []domain.Tweet{
						{TweetID: "tweet1", Content: "Hello World", UserID: "user1"},
					},
				},
			},
			want: want{
				err: false,
			},
		},
		{
			name: "Update with empty timeline",
			input: input{
				initialData: map[string]domain.Timeline{
					"user1": {
						UserID: "user1",
						Tweets: []domain.Tweet{
							{TweetID: "tweet1", Content: "Hello World", UserID: "user1"},
						},
					},
				},
				timeline: domain.Timeline{
					UserID: "user1",
					Tweets: []domain.Tweet{},
				},
			},
			want: want{
				err: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &InMemoryTimelineRepository{
				timelines: make(map[string]domain.Timeline),
			}

			for k, v := range tt.input.initialData {
				repo.timelines[k] = v
			}

			err := repo.Update(context.Background(), tt.input.timeline)

			if tt.want.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				storedTimeline, exists := repo.timelines[tt.input.timeline.UserID]
				assert.True(t, exists)
				assert.Equal(t, tt.input.timeline, storedTimeline)
			}
		})
	}
}

func TestNewInMemoryTimelineRepository(t *testing.T) {
	repo := NewInMemoryTimelineRepository()

	assert.NotNil(t, repo)

	timeline := domain.Timeline{
		UserID: "test-user",
		Tweets: []domain.Tweet{
			{TweetID: "test-tweet", Content: "Test Content", UserID: "test-user"},
		},
	}

	err := repo.Update(context.Background(), timeline)
	assert.NoError(t, err)

	retrievedTimeline, err := repo.Get(context.Background(), "test-user")
	assert.NoError(t, err)
	assert.Equal(t, timeline, retrievedTimeline)
}
