// In the "inmemory" package, create a file named "query_bus.go" for the in-memory query bus.

package inmemory

import (
	"context"
	"mail_service/internal/kit/query"
)

// QueryBus is an in-memory implementation of the query.Bus.
type QueryBus struct {
	handlers map[query.Type]query.Handler
}

// NewQueryBus initializes a new instance of QueryBus.
func NewQueryBus() *QueryBus {
	return &QueryBus{
		handlers: make(map[query.Type]query.Handler),
	}
}

// Dispatch implements the query.Bus interface.
func (b *QueryBus) Dispatch(ctx context.Context, qry query.Query) (interface{}, error) {
	handler, ok := b.handlers[qry.Type()]
	if !ok {
		return nil, nil // You can handle the case where no handler is registered for the query type.
	}

	return handler.Handle(ctx, qry)
}

// Register implements the query.Bus interface.
func (b *QueryBus) Register(qryType query.Type, handler query.Handler) {
	b.handlers[qryType] = handler
}
