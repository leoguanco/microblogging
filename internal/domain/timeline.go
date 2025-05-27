package domain

type Timeline struct {
	UserID string
	Limit  int
	Total  int
	Tweets []Tweet
}
