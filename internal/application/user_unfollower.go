package application

import (
	"context"
	"microblogging/internal/domain/ports"
	"microblogging/pkg/logging"
)

type UserUnfollower struct {
	userRepository ports.UserRepository
	eventBus       ports.EventBus
}

func NewUserUnfollower(userRepository ports.UserRepository, eventBus ports.EventBus) *UserUnfollower {
	return &UserUnfollower{userRepository: userRepository, eventBus: eventBus}
}

func (u UserUnfollower) Unfollow(ctx context.Context, unFolloweeID, followerID string) error {
	user, err := u.userRepository.Get(ctx, followerID)
	if err != nil {
		return err
	}

	err = user.Unfollow(unFolloweeID)
	if err != nil {
		logging.GetLogger().WithError(err).
			WithField("follower_id", followerID).
			WithField("un_followee_id", unFolloweeID)
		return err
	}

	err = u.userRepository.Save(ctx, user)
	if err != nil {
		return err
	}

	err = u.eventBus.Publish(ctx, user.PullDomainEvents())
	return err
}
