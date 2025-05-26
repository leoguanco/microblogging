package domain

import "time"

type Tweet struct {
	TweetID   string
	UserID    string
	Content   string
	CreatedAt time.Time
}

func NewTweet(tweetID string, userID string, content string, createdAt time.Time) Tweet {
	return Tweet{
		TweetID:   tweetID,
		UserID:    userID,
		Content:   content,
		CreatedAt: createdAt,
	}
}
