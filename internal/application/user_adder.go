package application

import (
	"context"
	"microblogging/internal/domain"
	"microblogging/internal/domain/ports"
)

type UserAdder struct {
	userRepository ports.UserRepository
}

func NewUserAdder(userRepository ports.UserRepository) *UserAdder {
	return &UserAdder{
		userRepository: userRepository,
	}
}

func (u *UserAdder) AddUser(ctx context.Context, userID string) error {
	user := domain.User{
		UserID: userID,
	}

	return u.userRepository.Save(ctx, user)
}
