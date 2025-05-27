package domain

type Event interface {
	GetType() string
	GetAggregateID() string
}
