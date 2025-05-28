package application

import (
	"context"
	"microblogging/internal/domain"
	"microblogging/pkg/logging"
)

type TimelineHandler struct {
	timelineUpdater *TimelineUpdater
}

func NewTimelineHandler(updater *TimelineUpdater) *TimelineHandler {
	return &TimelineHandler{timelineUpdater: updater}
}

func (h TimelineHandler) Handle(ctx context.Context, event domain.Event) error {
	switch e := event.(type) {
	case domain.TweetCreatedDomainEvent:
		tweet := domain.Tweet{TweetID: e.TweetID, UserID: e.UserID, Content: e.Content, CreatedAt: e.CreatedAt}
		return h.timelineUpdater.Update(ctx, e.UserID, tweet)
	}

	logging.GetLogger().WithField("event", event.GetType()).Warn("event not handled")
	return nil
}
