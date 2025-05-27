package ports

import (
	"context"
	"microblogging/internal/domain"
)

//go:generate mockgen -destination=mocks/timeline_repository_mock.go -package=mocks . TimelineRepository
type TimelineRepository interface {
	Get(ctx context.Context, userID string) (domain.Timeline, error)
}
