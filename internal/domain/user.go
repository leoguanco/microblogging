package domain

import "errors"

type User struct {
	UserID       string
	Follows      []string
	domainEvents []Event
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
