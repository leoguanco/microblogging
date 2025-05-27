package application

import "microblogging/internal/domain"

type TimelineHandler struct {
	timelineUpdater *TimelineUpdater
}

func NewTimelineHandler(updater *TimelineUpdater) *TimelineHandler {
	return &TimelineHandler{timelineUpdater: updater}
}

func (h TimelineHandler) Handle(event domain.Event) error {
	return nil
}
