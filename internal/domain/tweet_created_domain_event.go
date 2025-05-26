package domain

import "time"

func NewTweetCreatedDomainEvent(tweetID, userID, content string, createdAt time.Time) Event {
	return Event{
		"AggregateID": tweetID,
		"Type":        "tweet_created",
		"UserID":      userID,
		"Content":     content,
		"CreatedAt":   createdAt,
	}
}
