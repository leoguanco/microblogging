package outgoing

import (
	"context"
	"fmt"
	"microblogging/internal/domain/ports"
	"sync"
)

type InMemoryFolloweeRepository struct {
	followees map[string][]string
	mu        sync.RWMutex
}

func NewInMemoryFolloweeRepository() ports.FolloweeRepository {
	return &InMemoryFolloweeRepository{
		followees: make(map[string][]string),
	}
}

func (r *InMemoryFolloweeRepository) Get(ctx context.Context, followeeID string) ([]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	followers, exists := r.followees[followeeID]
	if !exists {
		return []string{}, fmt.Errorf("user with ID %s not found", followeeID)
	}
	return followers, nil
}

func (r *InMemoryFolloweeRepository) Save(ctx context.Context, followeeID string, followers []string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.followees[followeeID] = followers
	return nil
}
