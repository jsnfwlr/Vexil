package handlers

import (
	"context"
	"net/http"

	otelTrace "go.opentelemetry.io/otel/trace"

	"github.com/jsnfwlr/o11y"
	"github.com/jsnfwlr/vexil/internal/api/oapi"
	"github.com/jsnfwlr/vexil/internal/db"
	"github.com/jsnfwlr/vexil/internal/log"
)

func doBulkFlags(ctx context.Context, r oapi.BulkFlagsRequestObject) (res oapi.BulkFlagsResponseObject, fault error) {
	return oapi.BulkFlags200JSONResponse{}, nil
}

func doOptionsFlag(ctx context.Context, r oapi.OptionsFlagRequestObject) (res oapi.OptionsFlagResponseObject, fault error) {
	return oapi.OptionsFlag200Response{}, nil
}

type CreateFlagQueryProvider interface {
	AddFlag(ctx context.Context, arg db.AddFlagParams) (db.FeatureFlag, error)
	SetEnvFlagToDefault(ctx context.Context, arg db.SetEnvFlagToDefaultParams) error
}

func doCreateFlag(ctx context.Context, q CreateFlagQueryProvider, r oapi.CreateFlagRequestObject) (res oapi.CreateFlagResponseObject, fault error) {
	ctx, span := tracer.Start(ctx, "doCreateFlag", otelTrace.WithSpanKind(otelTrace.SpanKindServer))
	defer span.End()
	o := o11y.Get(ctx)

	o.Info("adding flag", span, "flag_name", r.Body.Name)

	addParams := db.AddFlagParams{
		Name:         r.Body.Name,
		DefaultValue: r.Body.DefaultValue,
		ValueType:    db.FlagType(r.Body.Type),
	}

	f, err := q.AddFlag(ctx, addParams)
	if err != nil {
		o.Error(err, span, log.EnvironmentNameKey, r.Body.Name)
		return oapi.CreateFlag500JSONResponse{
			Error:     err.Error(),
			ErrorCode: http.StatusInternalServerError,
		}, err
	}

	valueParams := db.SetEnvFlagToDefaultParams{
		FlagID:       f.ID,
		DefaultValue: r.Body.DefaultValue,
	}

	err = q.SetEnvFlagToDefault(ctx, valueParams)
	if err != nil {
		o.Error(err, span, log.EnvironmentNameKey, r.Body.Name)
		return oapi.CreateFlag500JSONResponse{
			Error:     err.Error(),
			ErrorCode: http.StatusInternalServerError,
		}, err
	}

	return oapi.CreateFlag201JSONResponse{}, nil
}

func doDeleteFlag(ctx context.Context, r oapi.DeleteFlagRequestObject) (res oapi.DeleteFlagResponseObject, fault error) {
	ctx, span := tracer.Start(ctx, "doDeleteFlag", otelTrace.WithSpanKind(otelTrace.SpanKindServer))
	defer span.End()

	return oapi.DeleteFlag204Response{}, nil
}

func doOptionsFlagId(ctx context.Context, r oapi.OptionsFlagIdRequestObject) (res oapi.OptionsFlagIdResponseObject, fault error) {
	ctx, span := tracer.Start(ctx, "doOptionsFlagId", otelTrace.WithSpanKind(otelTrace.SpanKindServer))
	defer span.End()

	return oapi.OptionsFlagId200Response{}, nil
}

func doUpdateFlag(ctx context.Context, r oapi.UpdateFlagRequestObject) (res oapi.UpdateFlagResponseObject, fault error) {
	ctx, span := tracer.Start(ctx, "doUpdateFlag", otelTrace.WithSpanKind(otelTrace.SpanKindServer))
	defer span.End()

	return oapi.UpdateFlag200JSONResponse{}, nil
}

func doOptionsEnvironmentNameFlagId(ctx context.Context, r oapi.OptionsEnvironmentNameFlagIdRequestObject) (res oapi.OptionsEnvironmentNameFlagIdResponseObject, fault error) {
	ctx, span := tracer.Start(ctx, "doOptionsEnvironmentNameFlagId", otelTrace.WithSpanKind(otelTrace.SpanKindServer))
	defer span.End()

	return oapi.OptionsEnvironmentNameFlagId200Response{}, nil
}

func doUpdateEnvironmentFlagValue(ctx context.Context, r oapi.UpdateEnvironmentFlagValueRequestObject) (res oapi.UpdateEnvironmentFlagValueResponseObject, fault error) {
	ctx, span := tracer.Start(ctx, "doUpdateEnvironmentFlagValue", otelTrace.WithSpanKind(otelTrace.SpanKindServer))
	defer span.End()

	return oapi.UpdateEnvironmentFlagValue200JSONResponse{}, nil
}
