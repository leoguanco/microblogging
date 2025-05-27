package usecases

import "context"

//go:generate mockgen -destination=mocks/user_unfollower_mock.go -package=mocks . UserUnfollower
type UserUnfollower interface {
	Unfollow(ctx context.Context, unFolloweeID, followerID string) error
}
