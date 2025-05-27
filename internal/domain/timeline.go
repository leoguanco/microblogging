package domain

type Timeline struct {
	UserID string
	Limit  int
	Total  int
	Tweets []Tweet
}

func (t *Timeline) AddTweet(tweet Tweet) {
	t.Tweets = append(t.Tweets, tweet)
	t.Total++
}
