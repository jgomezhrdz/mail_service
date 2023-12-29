package inmemory

import (
	"context"
	"fmt"
	"mail_service/internal/kit/event"
	"sync"
)

// EventBus is an in-memory implementation of the event.Bus.
type EventBus struct {
	handlers sync.Map
}

// NewEventBus initializes a new EventBus.
func NewEventBus() *EventBus {
	return &EventBus{}
}

// Publish implements the event.Bus interface.
func (b *EventBus) Publish(ctx context.Context, events []event.Event) error {
	for _, evt := range events {
		handlers, ok := b.handlers.Load(evt.Type())
		if !ok {
			continue
		}

		var wg sync.WaitGroup
		for _, handler := range handlers.([]event.Handler) {
			wg.Add(1)
			go func(wg *sync.WaitGroup, h event.Handler, e event.Event) {
				defer wg.Done()
				if err := h.Handle(ctx, e); err != nil {
					// Handle or log the error as needed.
					fmt.Printf("Error handling event: %v\n", err)
				}
			}(&wg, handler, evt)
		}

		// Wait for all goroutines to finish before processing the next event.
		wg.Wait()
	}

	return nil
}

// Subscribe implements the event.Bus interface.
func (b *EventBus) Subscribe(evtType event.Type, handler event.Handler) {
	val, loaded := b.handlers.LoadOrStore(evtType, []event.Handler{handler})
	if loaded {
		b.handlers.Store(evtType, append(val.([]event.Handler), handler))
	}
}
