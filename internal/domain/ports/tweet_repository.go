package ports

import (
	"context"
	"microblogging/internal/domain"
)

//go:generate mockgen -destination=mocks/tweet_repository_mock.go -package=mocks . TweetRepository
type TweetRepository interface {
	Save(ctx context.Context, tweet domain.Tweet) error
}
