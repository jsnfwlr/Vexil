package handlers

import (
	"context"

	"go.opentelemetry.io/otel"
	otelTrace "go.opentelemetry.io/otel/trace"

	"github.com/jsnfwlr/o11y"

	"github.com/jsnfwlr/vexil/internal/api/oapi"
	"github.com/jsnfwlr/vexil/internal/log"
)

type HealthCheckQueryProvider interface{}

var tracer = otel.Tracer("github.com/jsnfwlr/vexil/internal/api/handlers")

// doHealthCheck handles the GET request for the health check endpoint
func doHealthCheck(ctx context.Context, r oapi.HealthCheckRequestObject) (res oapi.HealthCheckResponseObject, fault error) {
	ctx, span := tracer.Start(ctx, "doHealthCheck", otelTrace.WithSpanKind(otelTrace.SpanKindServer))
	defer span.End()

	o := o11y.Get(ctx)
	o.Debug("healthCheck request", span, log.RequestIdKey, o11y.GetRequestID(ctx))

	// @TODO @jsnfwlr: Implement health check logic
	// This could include checking database connections, service availability, etc.
	// For now, we will just return a successful response.

	return oapi.HealthCheck200Response{}, nil
}

// doOptionsHealthCheck handles the OPTIONS request for the health check endpoint
func doOptionsHealthCheck(ctx context.Context, r oapi.OptionsHealthCheckRequestObject) (res oapi.OptionsHealthCheckResponseObject, fault error) {
	ctx, span := tracer.Start(ctx, "doOptionsHealthCheck", otelTrace.WithSpanKind(otelTrace.SpanKindServer))
	defer span.End()

	o := o11y.Get(ctx)
	o.Debug("optionsHealthCheck request", span, log.RequestIdKey, o11y.GetRequestID(ctx))

	return oapi.OptionsHealthCheck200Response{}, nil
}
