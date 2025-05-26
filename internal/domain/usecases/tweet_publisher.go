package usecases

import (
	"context"
	"microblogging/internal/domain"
)

//go:generate mockgen -destination=mocks/tweet_publisher_mock.go -package=mocks . TweetPublisher
type TweetPublisher interface {
	Publish(ctx context.Context, tweet domain.Tweet) error
}
