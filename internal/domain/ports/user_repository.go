package ports

import (
	"context"
	"microblogging/internal/domain"
)

//go:generate mockgen -destination=mocks/user_repository_mock.go -package=mocks . UserRepository
type UserRepository interface {
	Get(ctx context.Context, userID string) (domain.User, error)
	Save(ctx context.Context, user domain.User) error
}
