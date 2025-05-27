package outgoing

import (
	"context"
	"microblogging/internal/domain"
	"microblogging/internal/domain/ports"
	"sync"
)

type InMemoryTimelineRepository struct {
	timelines map[string]domain.Timeline
	mu        sync.RWMutex
}

func (r *InMemoryTimelineRepository) Update(ctx context.Context, timeline domain.Timeline) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.timelines[timeline.UserID] = timeline
	return nil
}

func NewInMemoryTimelineRepository() ports.TimelineRepository {
	return &InMemoryTimelineRepository{
		timelines: make(map[string]domain.Timeline),
	}
}

func (r *InMemoryTimelineRepository) Get(ctx context.Context, userID string) (domain.Timeline, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	timeline, exists := r.timelines[userID]
	if !exists {
		return domain.Timeline{UserID: userID, Tweets: []domain.Tweet{}}, nil
	}
	return timeline, nil
}
