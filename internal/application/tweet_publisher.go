package application

import (
	"context"
	"microblogging/internal/domain"
	"microblogging/internal/domain/ports"
)

type TweetPublisher struct {
	tweetRepository ports.TweetRepository
}

func NewTweetPublisher(tweetRepository ports.TweetRepository) *TweetPublisher {
	return &TweetPublisher{tweetRepository: tweetRepository}
}

func (p TweetPublisher) Publish(ctx context.Context, tweet domain.Tweet) error {
	err := p.tweetRepository.Save(ctx, tweet)

	return err
}
