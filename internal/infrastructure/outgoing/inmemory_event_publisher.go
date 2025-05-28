package outgoing

import (
	"context"
	"microblogging/internal/domain"
	"microblogging/internal/domain/usecases"
	"sync"
)

type InMemoryEventPublisher struct {
	handlers []usecases.EventHandler
	mutex    sync.RWMutex
}

func NewInMemoryEventPublisher() *InMemoryEventPublisher {
	return &InMemoryEventPublisher{
		handlers: make([]usecases.EventHandler, 0),
	}
}

func (p *InMemoryEventPublisher) Publish(ctx context.Context, events []domain.Event) error {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	for _, event := range events {
		for _, handler := range p.handlers {
			if err := handler.Handle(ctx, event); err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *InMemoryEventPublisher) RegisterHandler(handler usecases.EventHandler) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.handlers = append(p.handlers, handler)
}
