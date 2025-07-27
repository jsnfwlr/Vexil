package handlers

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"github.com/jsnfwlr/o11y"

	"github.com/jsnfwlr/vexil/internal/api/oapi"
	"github.com/jsnfwlr/vexil/internal/log"
)

type HealthCheckQueryProvider interface{}

var tracer = otel.Tracer("github.com/jsnfwlr/vexil/internal/api/handlers")

// doHealthCheck handles the GET request for the health check endpoint
func doHealthCheck(ctx context.Context, r oapi.HealthCheckRequestObject) (res oapi.HealthCheckResponseObject, fault error) {
	ctx, o := o11y.Get(tracer.Start(ctx, "HealthCheck", trace.WithSpanKind(trace.SpanKindServer)))
	defer o.End()

	o.Debug("healthCheck request", log.RequestIdKey, o11y.GetRequestID(ctx))

	// @TODO @jsnfwlr: Implement health check logic
	// This could include checking database connections, service availability, etc.
	// For now, we will just return a successful response.

	return oapi.HealthCheck200Response{}, nil
}

// doOptionsHealthCheck handles the OPTIONS request for the health check endpoint
func doOptionsHealthCheck(ctx context.Context, r oapi.OptionsHealthCheckRequestObject) (res oapi.OptionsHealthCheckResponseObject, fault error) {
	ctx, o := o11y.Get(tracer.Start(ctx, "OptionsHealthCheck", trace.WithSpanKind(trace.SpanKindServer)))
	defer o.End()

	o.Debug("optionsHealthCheck request", log.RequestIdKey, o11y.GetRequestID(ctx))

	return oapi.OptionsHealthCheck200Response{}, nil
}
