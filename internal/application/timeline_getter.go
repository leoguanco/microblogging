package application

import (
	"context"
	"microblogging/internal/domain"
	"microblogging/internal/domain/ports"
)

type TimelineGetter struct {
	timelineRepository ports.TimelineRepository
}

func NewTimelineGetter(timelineRepository ports.TimelineRepository) *TimelineGetter {
	return &TimelineGetter{timelineRepository: timelineRepository}
}

func (g TimelineGetter) Get(ctx context.Context, userID string) (domain.Timeline, error) {
	return g.timelineRepository.Get(ctx, userID)
}
