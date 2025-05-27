package domain

func NewUserUnfollowedDomainEvent(userID, unFolloweeID string) Event {
	return Event{
		"AggregateID": userID,
		"Type":        "user_unfollowed",
		"FolloweeID":  unFolloweeID,
	}
}
