package handlers

import (
	"context"

	"github.com/jsnfwlr/vexil/internal/api/oapi"
	otelTrace "go.opentelemetry.io/otel/trace"
)

func doListEnvironments(ctx context.Context, r oapi.ListEnvironmentsRequestObject) (res oapi.ListEnvironmentsResponseObject, fault error) {
	ctx, span := tracer.Start(ctx, "doListEnvironments", otelTrace.WithSpanKind(otelTrace.SpanKindServer))
	defer span.End()

	return oapi.ListEnvironments200JSONResponse{}, nil
}

func doOptionsEnvironment(ctx context.Context, r oapi.OptionsEnvironmentRequestObject) (res oapi.OptionsEnvironmentResponseObject, fault error) {
	ctx, span := tracer.Start(ctx, "doOptionsEnvironment", otelTrace.WithSpanKind(otelTrace.SpanKindServer))
	defer span.End()

	return oapi.OptionsEnvironment200Response{}, nil
}

func doCreateEnvironment(ctx context.Context, r oapi.CreateEnvironmentRequestObject) (res oapi.CreateEnvironmentResponseObject, fault error) {
	ctx, span := tracer.Start(ctx, "doCreateEnvironment", otelTrace.WithSpanKind(otelTrace.SpanKindServer))
	defer span.End()

	return oapi.CreateEnvironment201JSONResponse{}, nil
}

func doDeleteEnvironment(ctx context.Context, r oapi.DeleteEnvironmentRequestObject) (res oapi.DeleteEnvironmentResponseObject, fault error) {
	ctx, span := tracer.Start(ctx, "doDeleteEnvironment", otelTrace.WithSpanKind(otelTrace.SpanKindServer))
	defer span.End()

	return oapi.DeleteEnvironment204Response{}, nil
}

func doOptionsEnvironmentName(ctx context.Context, r oapi.OptionsEnvironmentNameRequestObject) (res oapi.OptionsEnvironmentNameResponseObject, fault error) {
	ctx, span := tracer.Start(ctx, "doOptionsEnvironmentName", otelTrace.WithSpanKind(otelTrace.SpanKindServer))
	defer span.End()

	return oapi.OptionsEnvironmentName200Response{}, nil
}
