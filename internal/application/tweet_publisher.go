package application

import (
	"context"
	"microblogging/internal/domain"
	"microblogging/internal/domain/ports"
	"time"
)

type TweetCreator struct {
	tweetRepository ports.TweetRepository
	eventBus        ports.EventBus
}

func NewTweetPublisher(tweetRepository ports.TweetRepository, eventBus ports.EventBus) *TweetCreator {
	return &TweetCreator{tweetRepository: tweetRepository, eventBus: eventBus}
}

func (p TweetCreator) Create(ctx context.Context, tweetID, userID, content string, createdAt time.Time) error {
	tweet, err := domain.NewTweet(tweetID, userID, content, createdAt)
	if err != nil {
		return err
	}

	err = p.tweetRepository.Save(ctx, *tweet)
	if err != nil {
		return err
	}

	err = p.eventBus.Publish(ctx, tweet.PullDomainEvents())

	return nil
}
