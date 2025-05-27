package usecases

import "microblogging/internal/domain"

//go:generate mockgen -destination=mocks/event_handler_mock.go -package=mocks . EventHandler
type EventHandler interface {
	Handle(event domain.Event) error
}
