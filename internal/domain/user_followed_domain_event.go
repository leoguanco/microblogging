package domain

type UserFollowedDomainEvent struct {
	UserID     string
	FolloweeID string
}

func (e UserFollowedDomainEvent) GetType() string {
	return "user_followed"
}

func (e UserFollowedDomainEvent) GetAggregateID() string {
	return e.UserID
}

func NewUserFollowedDomainEvent(userID, followeeID string) UserFollowedDomainEvent {
	return UserFollowedDomainEvent{
		UserID:     userID,
		FolloweeID: followeeID,
	}
}
