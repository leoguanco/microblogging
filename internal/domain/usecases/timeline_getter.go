package usecases

import (
	"context"
	"microblogging/internal/domain"
)

//go:generate mockgen -destination=mocks/timeline_getter_mock.go -package=mocks . TimelineGetter
type TimelineGetter interface {
	Get(ctx context.Context, userID string) (domain.Timeline, error)
}
