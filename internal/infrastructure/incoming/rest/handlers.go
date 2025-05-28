package rest

import (
	"microblogging/internal/domain/usecases"
	"microblogging/pkg/logging"
)

type Handlers struct {
	TweetHandler    *TweetHandler
	UserHandler     *UserHandler
	TimelineHandler *TimelineHandler
	Logger          *logging.Logger
}

func NewHandlers(
	tweetCreator usecases.TweetCreator,
	userFollower usecases.UserFollower,
	userUnfollower usecases.UserUnfollower,
	timelineGetter usecases.TimelineGetter,
	userAdder usecases.UserAdder,
) *Handlers {
	logger := logging.GetLogger().WithField("component", "APIHandlers")

	return &Handlers{
		TweetHandler:    NewTweetHandler(tweetCreator),
		UserHandler:     NewUserHandler(userFollower, userUnfollower, userAdder),
		TimelineHandler: NewTimelineHandler(timelineGetter),
		Logger:          logger,
	}
}
