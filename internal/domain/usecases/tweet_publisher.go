package usecases

import (
	"context"
)

//go:generate mockgen -destination=mocks/tweet_creator_mock.go -package=mocks . TweetCreator
type TweetCreator interface {
	Create(ctx context.Context, tweetID, userID, content string) error
}
