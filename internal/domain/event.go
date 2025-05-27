package domain

import "time"

type Event interface {
	GetType() string
	GetAggregateID() string
	GetTimestamp() time.Time
}
