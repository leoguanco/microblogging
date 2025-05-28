package domain

import "errors"

type User struct {
	UserID       string
	Follows      []string
	domainEvents []Event
	Followers    []string
}

func (u *User) PullDomainEvents() []Event {
	events := u.domainEvents
	u.domainEvents = []Event{}

	return events
}

func (u *User) Record(event Event) {
	u.domainEvents = append(u.domainEvents, event)
}

func (u *User) Follow(followeeID string) error {
	if u.UserID == followeeID {
		return errors.New("cannot follow yourself")
	}

	if u.AlreadyFollow(followeeID) {
		return errors.New("already follow")
	}

	u.Follows = append(u.Follows, followeeID)
	u.Record(NewUserFollowedDomainEvent(u.UserID, followeeID))

	return nil
}

func (u *User) AlreadyFollow(id string) bool {
	for _, followID := range u.Follows {
		if followID == id {
			return true
		}
	}

	return false
}

func (u *User) Unfollow(unFolloweeID string) error {
	if u.UserID == unFolloweeID {
		return errors.New("cannot unfollow yourself")
	}

	var follows []string
	for _, followID := range u.Follows {
		if followID != unFolloweeID {
			follows = append(follows, followID)
		}
	}
	u.Follows = follows
	u.Record(NewUserUnfollowedDomainEvent(u.UserID, unFolloweeID))

	return nil
}

func (u *User) GetFollowers() []string {
	return u.Followers
}

func (u *User) AddFollower(followerID string) {
	u.Followers = append(u.Followers, followerID)
}

func (u *User) RemoveFollower(followerID string) {
	var followers []string
	for _, fID := range u.Followers {
		if fID != followerID {
			followers = append(followers, fID)
		}
	}
	u.Followers = followers
}
