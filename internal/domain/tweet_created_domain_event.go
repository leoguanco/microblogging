package domain

import "time"

type TweetCreatedDomainEvent struct {
	TweetID   string
	UserID    string
	Content   string
	CreatedAt time.Time
}

func (e TweetCreatedDomainEvent) GetType() string {
	return "tweet_created"
}

func (e TweetCreatedDomainEvent) GetAggregateID() string {
	return e.TweetID
}

func NewTweetCreatedDomainEvent(tweetID, userID, content string, createdAt time.Time) TweetCreatedDomainEvent {
	return TweetCreatedDomainEvent{
		TweetID:   tweetID,
		UserID:    userID,
		Content:   content,
		CreatedAt: createdAt,
	}
}
