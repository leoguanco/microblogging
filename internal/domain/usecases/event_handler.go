package usecases

import (
	"context"
	"microblogging/internal/domain"
)

//go:generate mockgen -destination=mocks/event_handler_mock.go -package=mocks . EventHandler
type EventHandler interface {
	Handle(ctx context.Context, event domain.Event) error
}
