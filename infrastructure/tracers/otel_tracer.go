package tracers

import (
	"context"
	"presentation-advert-read-api/application/tracers"
)

type exampleTracer struct{}

func NewExampleTracer() tracers.Tracer {
	return &exampleTracer{}
}

func (tc *exampleTracer) Trace(ctx context.Context, structName, funcName string) (context.Context, tracers.DeferFunc) {
	//span := trace.SpanFromContext(ctx)
	//if span == nil || !span.SpanContext().IsValid() {
	//	return ctx, func() {
	//		// do nothing
	//	}
	//}
	//tracer := span.TracerProvider().Tracer(structName)
	//childCtx, childSpan := tracer.Start(ctx, funcName)
	//return childCtx, func() {
	//	childSpan.End()
	//}
	return ctx, func() {}
}
