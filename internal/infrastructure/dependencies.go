package infrastructure

import (
	"microblogging/internal/application"
	"microblogging/internal/domain/ports"
	"microblogging/internal/domain/usecases"
	"microblogging/internal/infrastructure/outgoing"
)

type AppDependencies struct {
	TweetCreator   usecases.TweetCreator
	UserFollower   usecases.UserFollower
	UserUnfollower usecases.UserUnfollower
	TimelineGetter usecases.TimelineGetter
}

func BuildDependencies() AppDependencies {
	userRepo := outgoing.NewInMemoryUserRepository()
	tweetRepo := outgoing.NewInMemoryTweetRepository()
	timelineRepo := outgoing.NewInMemoryTimelineRepository()

	timelineUpdater := application.NewTimelineUpdater(timelineRepo)
	timelineUpdaterHandler := application.NewTimelineHandler(timelineUpdater)

	eventPublisher := outgoing.NewInMemoryEventPublisher()
	eventPublisher.RegisterHandler(timelineUpdaterHandler)

	tweetCreator := application.NewTweetCreator(tweetRepo, eventPublisher)
	userFollower := application.NewUserFollower(userRepo, eventPublisher)
	userUnfollower := application.NewUserUnfollower(userRepo, eventPublisher)
	timelineGetter := application.NewTimelineGetter(timelineRepo)

	return AppDependencies{
		TweetCreator:   tweetCreator,
		UserFollower:   userFollower,
		UserUnfollower: userUnfollower,
		TimelineGetter: timelineGetter,
	}
}

func CreateRepositories() (ports.UserRepository, ports.TweetRepository, ports.TimelineRepository) {
	userRepo := outgoing.NewInMemoryUserRepository()
	tweetRepo := outgoing.NewInMemoryTweetRepository()
	timelineRepo := outgoing.NewInMemoryTimelineRepository()

	return userRepo, tweetRepo, timelineRepo
}
