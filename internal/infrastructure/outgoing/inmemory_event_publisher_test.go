package outgoing

import (
	"context"
	"errors"
	"microblogging/internal/domain"
	"microblogging/internal/domain/usecases"
	"testing"
)

type mockEventHandler struct {
	handleFunc func(ctx context.Context, event domain.Event) error
}

func (m *mockEventHandler) Handle(ctx context.Context, event domain.Event) error {
	return m.handleFunc(ctx, event)
}

func TestInMemoryEventPublisher_Publish(t *testing.T) {
	tests := []struct {
		name     string
		handlers []usecases.EventHandler
		events   []domain.Event
		wantErr  bool
	}{
		{
			name:     "no handlers, no events",
			handlers: nil,
			events:   nil,
			wantErr:  false,
		},
		{
			name: "no events, one handler",
			handlers: []usecases.EventHandler{
				&mockEventHandler{
					handleFunc: func(ctx context.Context, event domain.Event) error {
						return nil
					},
				},
			},
			events:  nil,
			wantErr: false,
		},
		{
			name: "one handler, successful handling",
			handlers: []usecases.EventHandler{
				&mockEventHandler{
					handleFunc: func(ctx context.Context, event domain.Event) error {
						return nil
					},
				},
			},
			events:  []domain.Event{domain.TweetCreatedDomainEvent{TweetID: "1"}, domain.TweetCreatedDomainEvent{TweetID: "2"}},
			wantErr: false,
		},
		{
			name: "one handler, handler error",
			handlers: []usecases.EventHandler{
				&mockEventHandler{
					handleFunc: func(ctx context.Context, event domain.Event) error {
						return errors.New("handler error")
					},
				},
			},
			events:  []domain.Event{domain.TweetCreatedDomainEvent{TweetID: "1"}},
			wantErr: true,
		},
		{
			name: "multiple handlers, successful handling",
			handlers: []usecases.EventHandler{
				&mockEventHandler{
					handleFunc: func(ctx context.Context, event domain.Event) error {
						return nil
					},
				},
				&mockEventHandler{
					handleFunc: func(ctx context.Context, event domain.Event) error {
						return nil
					},
				},
			},
			events:  []domain.Event{domain.TweetCreatedDomainEvent{TweetID: "1"}, domain.TweetCreatedDomainEvent{TweetID: "2"}},
			wantErr: false,
		},
		{
			name: "multiple handlers, one handler error",
			handlers: []usecases.EventHandler{
				&mockEventHandler{
					handleFunc: func(ctx context.Context, event domain.Event) error {
						return nil
					},
				},
				&mockEventHandler{
					handleFunc: func(ctx context.Context, event domain.Event) error {
						return errors.New("handler error")
					},
				},
			},
			events:  []domain.Event{domain.TweetCreatedDomainEvent{TweetID: "1"}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			publisher := NewInMemoryEventPublisher()
			for _, handler := range tt.handlers {
				publisher.RegisterHandler(handler)
			}

			err := publisher.Publish(context.Background(), tt.events)
			if (err != nil) != tt.wantErr {
				t.Errorf("Publish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInMemoryEventPublisher_RegisterHandler(t *testing.T) {
	t.Run("register single handler", func(t *testing.T) {
		publisher := NewInMemoryEventPublisher()
		handler := &mockEventHandler{
			handleFunc: func(ctx context.Context, event domain.Event) error {
				return nil
			},
		}

		publisher.RegisterHandler(handler)

		if len(publisher.handlers) != 1 {
			t.Errorf("expected 1 handler, got %d", len(publisher.handlers))
		}
	})

	t.Run("register multiple handlers", func(t *testing.T) {
		publisher := NewInMemoryEventPublisher()
		handler1 := &mockEventHandler{
			handleFunc: func(ctx context.Context, event domain.Event) error {
				return nil
			},
		}
		handler2 := &mockEventHandler{
			handleFunc: func(ctx context.Context, event domain.Event) error {
				return nil
			},
		}

		publisher.RegisterHandler(handler1)
		publisher.RegisterHandler(handler2)

		if len(publisher.handlers) != 2 {
			t.Errorf("expected 2 handlers, got %d", len(publisher.handlers))
		}
	})
}
