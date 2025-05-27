package outgoing

import (
	"context"
	"microblogging/internal/domain"
	"microblogging/internal/domain/ports"
	"sync"
)

type InMemoryTweetRepository struct {
	tweets map[string]domain.Tweet
	mu     sync.RWMutex
}

func NewInMemoryTweetRepository() ports.TweetRepository {
	return &InMemoryTweetRepository{
		tweets: make(map[string]domain.Tweet),
	}
}

func (r *InMemoryTweetRepository) Save(ctx context.Context, tweet domain.Tweet) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tweets[tweet.TweetID] = tweet
	return nil
}
