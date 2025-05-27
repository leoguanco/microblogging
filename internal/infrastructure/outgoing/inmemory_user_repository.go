package outgoing

import (
	"context"
	"fmt"
	"microblogging/internal/domain"
	"microblogging/internal/domain/ports"
	"sync"
)

type InMemoryUserRepository struct {
	users map[string]domain.User
	mu    sync.RWMutex
}

func NewInMemoryUserRepository() ports.UserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]domain.User),
	}
}

func (r *InMemoryUserRepository) Get(ctx context.Context, userID string) (domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[userID]
	if !exists {
		return domain.User{}, fmt.Errorf("user with ID %s not found", userID)
	}
	return user, nil
}

func (r *InMemoryUserRepository) Save(ctx context.Context, user domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[user.UserID] = user
	return nil
}
