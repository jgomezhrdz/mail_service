// query.go

package query

import "context"

// Type represents the type of a query.
type Type string

// Query is the interface that all queries must implement.
type Query interface {
	Type() Type
}

// Handler is the interface that query handlers must implement.
type Handler interface {
	Handle(context.Context, Query) (interface{}, error)
}

// Bus is the interface for a query bus.
type Bus interface {
	Dispatch(context.Context, Query) (interface{}, error)
	Register(Type, Handler)
}
