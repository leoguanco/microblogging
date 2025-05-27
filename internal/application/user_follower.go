package application

import (
	"context"
	"microblogging/internal/domain/ports"
	"microblogging/pkg/logging"
)

type UserFollower struct {
	userRepository ports.UserRepository
	eventBus       ports.EventBus
}

func NewUserFollower(userRepository ports.UserRepository, eventBus ports.EventBus) *UserFollower {
	return &UserFollower{userRepository: userRepository, eventBus: eventBus}
}

func (f UserFollower) Follow(ctx context.Context, followeeID, followerID string) error {
	user, err := f.userRepository.Get(ctx, followerID)
	if err != nil {
		return err
	}

	err = user.Follow(followeeID)
	if err != nil {
		logging.GetLogger().WithError(err).
			WithField("follower_id", followerID).
			WithField("followee_id", followeeID)
		return err
	}

	err = f.userRepository.Save(ctx, user)
	if err != nil {
		return err
	}

	err = f.eventBus.Publish(ctx, user.PullDomainEvents())
	return err
}
