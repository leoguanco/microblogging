package ports

import (
	"context"
	"microblogging/internal/domain"
)

//go:generate mockgen -destination=mocks/event_bus_mock.go -package=mocks . EventBus
type EventBus interface {
	Publish(ctx context.Context, event domain.Event) error
}
