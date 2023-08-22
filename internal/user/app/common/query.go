package common

import "context"

// QueryHandler is a query function that can be executed by the service.
type QueryHandler[Qry any, Rsp any] func(ctx context.Context, qry Qry) (Rsp, error)

// QueryDecorator is a function that wraps a query handler.
type QueryDecorator[Qry any, Rsp any] func(QueryHandler[Qry, Rsp]) QueryHandler[Qry, Rsp]

// ApplyQueryDecorators applies the given query decorators to the given query handler.
func ApplyQueryDecorators[Qry any, Rsp any](handler QueryHandler[Qry, Rsp], decorators ...QueryDecorator[Qry, Rsp]) QueryHandler[Qry, Rsp] {
	for _, decorator := range decorators {
		handler = decorator(handler)
	}
	return handler
}
