package ports

import "context"

//go:generate mockgen -destination=mocks/followee_repository_mock.go -package=mocks . FolloweeRepository
type FolloweeRepository interface {
	Get(ctx context.Context, followeeID string) ([]string, error)
	Save(ctx context.Context, followeeID string, followers []string) error
}
