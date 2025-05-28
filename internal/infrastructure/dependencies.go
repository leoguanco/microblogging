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
	UserAdder      usecases.UserAdder
}

func BuildDependencies() AppDependencies {
	userRepo := outgoing.NewInMemoryUserRepository()
	tweetRepo := outgoing.NewInMemoryTweetRepository()
	timelineRepo := outgoing.NewInMemoryTimelineRepository()
	followeeRepo := outgoing.NewInMemoryFolloweeRepository()

	timelineUpdater := application.NewTimelineUpdater(timelineRepo, followeeRepo)
	timelineUpdaterHandler := application.NewTimelineHandler(timelineUpdater)

	eventPublisher := outgoing.NewInMemoryEventPublisher()
	eventPublisher.RegisterHandler(timelineUpdaterHandler)

	tweetCreator := application.NewTweetCreator(tweetRepo, eventPublisher)
	userFollower := application.NewUserFollower(userRepo, followeeRepo, eventPublisher)
	userUnfollower := application.NewUserUnfollower(userRepo, eventPublisher)
	timelineGetter := application.NewTimelineGetter(timelineRepo)
	userAdder := application.NewUserAdder(userRepo)

	return AppDependencies{
		TweetCreator:   tweetCreator,
		UserFollower:   userFollower,
		UserUnfollower: userUnfollower,
		TimelineGetter: timelineGetter,
		UserAdder:      userAdder,
	}
}

func CreateRepositories() (ports.UserRepository, ports.TweetRepository, ports.TimelineRepository) {
	userRepo := outgoing.NewInMemoryUserRepository()
	tweetRepo := outgoing.NewInMemoryTweetRepository()
	timelineRepo := outgoing.NewInMemoryTimelineRepository()

	return userRepo, tweetRepo, timelineRepo
}
