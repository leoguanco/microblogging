package domain

type UserUnfollowedDomainEvent struct {
	UserID       string
	UnFolloweeID string
}

func (e UserUnfollowedDomainEvent) GetType() string {
	return "user_unfollowed"
}

func (e UserUnfollowedDomainEvent) GetAggregateID() string {
	return e.UserID
}

func NewUserUnfollowedDomainEvent(userID, unFolloweeID string) UserUnfollowedDomainEvent {
	return UserUnfollowedDomainEvent{
		UserID:       userID,
		UnFolloweeID: unFolloweeID,
	}
}
