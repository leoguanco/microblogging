package domain

func NewUserFollowedDomainEvent(userID, followeeID string) Event {
	return Event{
		"AggregateID": userID,
		"Type":        "user_followed",
		"FolloweeID":  followeeID,
	}
}
