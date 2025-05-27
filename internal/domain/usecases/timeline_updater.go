package usecases

import (
	"context"
	"microblogging/internal/domain"
)

//go:generate mockgen -destination=mocks/timeline_updater_mock.go -package=mocks . TimelineUpdater
type TimelineUpdater interface {
	Update(ctx context.Context, userID string, tweet domain.Tweet) error
}
