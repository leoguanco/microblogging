package application

import (
	"context"
	"microblogging/internal/domain/ports"
	"microblogging/pkg/logging"
)

type UserFollower struct {
	userRepository     ports.UserRepository
	followeeRepository ports.FolloweeRepository
	eventBus           ports.EventBus
}

func NewUserFollower(
	userRepository ports.UserRepository,
	followeeRepository ports.FolloweeRepository,
	eventBus ports.EventBus,
) *UserFollower {
	return &UserFollower{userRepository: userRepository, followeeRepository: followeeRepository, eventBus: eventBus}
}

func (f UserFollower) Follow(ctx context.Context, followeeID, followerID string) error {
	follower, err := f.userRepository.Get(ctx, followerID)
	if err != nil {
		return err
	}

	err = follower.Follow(followeeID)
	if err != nil {
		logging.GetLogger().WithError(err).
			WithField("follower_id", followerID).
			WithField("followee_id", followeeID)
		return err
	}

	err = f.userRepository.Save(ctx, follower)
	if err != nil {
		return err
	}

	followee, err := f.userRepository.Get(ctx, followeeID)
	if err != nil {
		return err
	}

	followee.AddFollower(followerID)
	err = f.followeeRepository.Save(ctx, followeeID, followee.GetFollowers())
	if err != nil {
		return err
	}

	err = f.userRepository.Save(ctx, followee)
	if err != nil {
		return err
	}

	err = f.eventBus.Publish(ctx, follower.PullDomainEvents())

	return err
}
