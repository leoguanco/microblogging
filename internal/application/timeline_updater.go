package application

import (
	"context"
	"microblogging/internal/domain"
	"microblogging/internal/domain/ports"
)

type TimelineUpdater struct {
	timelineRepository ports.TimelineRepository
}

func NewTimelineUpdater(timelineRepository ports.TimelineRepository) *TimelineUpdater {
	return &TimelineUpdater{timelineRepository: timelineRepository}
}

func (u TimelineUpdater) Update(ctx context.Context, userID string, tweet domain.Tweet) error {
	timeline, err := u.timelineRepository.Get(ctx, userID)
	if err != nil {
		return err
	}

	timeline.AddTweet(tweet)

	err = u.timelineRepository.Update(ctx, timeline)
	if err != nil {
		return err
	}

	return nil
}
