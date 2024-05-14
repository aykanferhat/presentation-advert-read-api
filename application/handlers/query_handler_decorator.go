package handlers

import (
	"context"
	"presentation-advert-read-api/application/tracers"
	"reflect"
)

type QueryHandlerInterface[T any, R any] interface {
	Handle(ctx context.Context, value T) (R, error)
}

type QueryHandlerDecorator[T any, R any] interface {
	Handle(ctx context.Context, value T) (R, error)
}

type queryHandlerDecorator[T any, R any] struct {
	tracers []tracers.Tracer
	handler QueryHandlerInterface[T, R]
}

func NewQueryHandlerDecorator[T any, R any](handler QueryHandlerInterface[T, R], tracers []tracers.Tracer) QueryHandlerDecorator[T, R] {
	return &queryHandlerDecorator[T, R]{
		tracers: tracers,
		handler: handler,
	}
}

func (decorator *queryHandlerDecorator[T, R]) Handle(ctx context.Context, value T) (R, error) {
	deferFunctions := make([]tracers.DeferFunc, 0, len(decorator.tracers))
	for _, decoratorTracer := range decorator.tracers {
		var deferFunction tracers.DeferFunc
		ctx, deferFunction = decoratorTracer.Trace(ctx, reflect.TypeOf(decorator.handler).String(), "Handle")
		deferFunctions = append(deferFunctions, deferFunction)
	}
	result, err := decorator.handler.Handle(ctx, value)
	for _, function := range deferFunctions {
		function()
	}
	return result, err
}
