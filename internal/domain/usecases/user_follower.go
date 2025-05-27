package usecases

import "context"

//go:generate mockgen -destination=mocks/user_follower_mock.go -package=mocks . UserFollower
type UserFollower interface {
	Follow(ctx context.Context, followeeID, followerID string) error
}
