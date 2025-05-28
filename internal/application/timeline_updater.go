package application

import (
	"context"
	"microblogging/internal/domain"
	"microblogging/internal/domain/ports"
)

type TimelineUpdater struct {
	timelineRepository ports.TimelineRepository
	followeeRepository ports.FolloweeRepository
}

func NewTimelineUpdater(timelineRepository ports.TimelineRepository, followeeRepository ports.FolloweeRepository) *TimelineUpdater {
	return &TimelineUpdater{timelineRepository: timelineRepository, followeeRepository: followeeRepository}
}

func (u TimelineUpdater) Update(ctx context.Context, userID string, tweet domain.Tweet) error {
	followers, err := u.followeeRepository.Get(ctx, userID)
	if err != nil {
		return err
	}

	for _, follower := range followers {
		timeline, err := u.timelineRepository.Get(ctx, follower)
		if err != nil {
			return err
		}

		timeline.AddTweet(tweet)

		err = u.timelineRepository.Update(ctx, timeline)
		if err != nil {
			return err
		}
	}
	return nil
}
