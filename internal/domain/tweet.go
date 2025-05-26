package domain

import (
	"errors"
	"microblogging/pkg/logging"
	"time"
)

const maxContent = 280

type Tweet struct {
	TweetID      string
	UserID       string
	Content      string
	CreatedAt    time.Time
	domainEvents []Event
}

func (t *Tweet) PullDomainEvents() []Event {
	events := t.domainEvents
	t.domainEvents = []Event{}

	return events
}

func (t *Tweet) Record(event Event) {
	t.domainEvents = append(t.domainEvents, event)
}

func NewTweet(tweetID string, userID string, content string, createdAt time.Time) (*Tweet, error) {
	if len(content) > maxContent {
		logging.GetLogger().WithField("content", content).Error("content is too long")
		return nil, errors.New("content is too long")
	}

	tweet := Tweet{
		TweetID:   tweetID,
		UserID:    userID,
		Content:   content,
		CreatedAt: createdAt,
	}

	tweet.Record(NewTweetCreatedDomainEvent(tweetID, userID, content, createdAt))

	return &tweet, nil
}
