package handlers

import (
	"context"
	"fmt"

	"github.com/jsnfwlr/o11y"
	"github.com/jsnfwlr/vexil/internal/api/oapi"
	"github.com/jsnfwlr/vexil/internal/db"
	"github.com/r3labs/sse/v2"
	otelTrace "go.opentelemetry.io/otel/trace"
)

type Handlers struct {
	DBClient   *db.Client
	EventSrv   *sse.Server
	staticPath string
	indexPath  string
}

type Flag struct {
	Name  string
	Value string
	Type  db.FlagType
	Env   string
}

func New(ctx context.Context, dbClient *db.Client, enableSSE bool, staticPath, indexPath string) (handlers Handlers, fault error) {
	var eventSrv *sse.Server
	ctx, span := tracer.Start(ctx, "New", otelTrace.WithSpanKind(otelTrace.SpanKindServer))
	defer span.End()

	o := o11y.Get(ctx)

	if enableSSE {
		eventSrv = sse.NewWithCallback(AddSub, RemSub)
		eventSrv.AutoReplay = false
		// @TODO - use the dbClient to query the db for all the environments, and create streams for each one
		envs, err := dbClient.Queries.ListEnvironments(ctx)
		if err != nil {
			return Handlers{}, err
		}
		for _, e := range envs {
			o.Info("creating environment stream", span, "environment", e)
			eventSrv.CreateStream(e)
		}
	}

	s := Handlers{
		DBClient:   dbClient,
		EventSrv:   eventSrv,
		staticPath: staticPath,
		indexPath:  indexPath,
	}
	return s, nil
}

func AddSub(streamId string, sub *sse.Subscriber) {
	fmt.Printf("Subscriber added to stream %s\n", streamId)
}

func RemSub(streamId string, sub *sse.Subscriber) {
	fmt.Printf("Subscriber left stream %s\n", streamId)
}

func (h Handlers) UIFindBulkFlags(ctx context.Context, r oapi.UIFindBulkFlagsRequestObject) (res oapi.UIFindBulkFlagsResponseObject, fault error) {
	return doUIFindBulkFlags(ctx, h.DBClient.Queries, r)
}

func (h Handlers) UIOptionsFindBulkFlags(ctx context.Context, r oapi.UIOptionsFindBulkFlagsRequestObject) (res oapi.UIOptionsFindBulkFlagsResponseObject, fault error) {
	return doUIOptionsFindFlags(ctx, r)
}

// HealthCheck checks the health of the server (GET /healthcheck)
func (h Handlers) HealthCheck(ctx context.Context, r oapi.HealthCheckRequestObject) (res oapi.HealthCheckResponseObject, fault error) {
	return doHealthCheck(ctx, r)
}

// OptionsHealthCheck checks the options for the endpoint (OPTIONS /healthcheck)
func (h Handlers) OptionsHealthCheck(ctx context.Context, r oapi.OptionsHealthCheckRequestObject) (res oapi.OptionsHealthCheckResponseObject, fault error) {
	return doOptionsHealthCheck(ctx, r)
}

// GetFlagsByEnvironment Get flags by environment (GET /api/environment/{environment_name}/flag)
func (h Handlers) GetFlagsByEnvironment(ctx context.Context, r oapi.GetFlagsByEnvironmentRequestObject) (res oapi.GetFlagsByEnvironmentResponseObject, fault error) {
	return doGetFlagsByEnvironment(ctx, h.DBClient.Queries, r)
}

// OptionsEnvironmentNameFlag Check the options for the endpoint (OPTIONS /api/environment/{environment_name}/flag)
func (h Handlers) OptionsEnvironmentNameFlag(ctx context.Context, r oapi.OptionsEnvironmentNameFlagRequestObject) (res oapi.OptionsEnvironmentNameFlagResponseObject, fault error) {
	return doOptionsEnvironmentNameFlag(ctx, r)
}

// BulkFlags Get bulk Find of flags (GET /api/flag)
func (h Handlers) BulkFlags(ctx context.Context, r oapi.BulkFlagsRequestObject) (res oapi.BulkFlagsResponseObject, fault error) {
	return doBulkFlags(ctx, r)
}

// OptionsFlag Check the options for the endpoint (OPTIONS /api/flag)
func (h Handlers) OptionsFlag(ctx context.Context, r oapi.OptionsFlagRequestObject) (res oapi.OptionsFlagResponseObject, fault error) {
	return doOptionsFlag(ctx, r)
}

// CreateFlag Create a new flag (POST /api/flag)
func (h Handlers) CreateFlag(ctx context.Context, r oapi.CreateFlagRequestObject) (res oapi.CreateFlagResponseObject, fault error) {
	return doCreateFlag(ctx, h.DBClient.Queries, r)
}

// DeleteFlag Delete a flag (DELETE /api/flag/{flag_id})
func (h Handlers) DeleteFlag(ctx context.Context, r oapi.DeleteFlagRequestObject) (res oapi.DeleteFlagResponseObject, fault error) {
	return doDeleteFlag(ctx, r)
}

// OptionsFlagId Check the options for the endpoint (OPTIONS /api/flag/{flag_id})
func (h Handlers) OptionsFlagId(ctx context.Context, r oapi.OptionsFlagIdRequestObject) (res oapi.OptionsFlagIdResponseObject, fault error) {
	return doOptionsFlagId(ctx, r)
}

// UpdateFlag Update a flag (PUT /api/flag/{flag_id})
func (h Handlers) UpdateFlag(ctx context.Context, r oapi.UpdateFlagRequestObject) (res oapi.UpdateFlagResponseObject, fault error) {
	return doUpdateFlag(ctx, r)
}

// OptionsEnvironmentNameFlagId Check the options for the endpoint (OPTIONS /api/environment/{environment_name}/flag/{flag_id})
func (h Handlers) OptionsEnvironmentNameFlagId(ctx context.Context, r oapi.OptionsEnvironmentNameFlagIdRequestObject) (res oapi.OptionsEnvironmentNameFlagIdResponseObject, fault error) {
	return doOptionsEnvironmentNameFlagId(ctx, r)
}

// UpdateEnvironmentFlagValue Update a flag in an environment (PUT /api/environment/{environment_name}/flag/{flag_id})
func (h Handlers) UpdateEnvironmentFlagValue(ctx context.Context, r oapi.UpdateEnvironmentFlagValueRequestObject) (res oapi.UpdateEnvironmentFlagValueResponseObject, fault error) {
	return doUpdateEnvironmentFlagValue(ctx, r)
}

// ListEnvironments Find environments (GET /api/environment)
func (h Handlers) ListEnvironments(ctx context.Context, r oapi.ListEnvironmentsRequestObject) (res oapi.ListEnvironmentsResponseObject, fault error) {
	return doListEnvironments(ctx, r)
}

// OptionsEnvironment Check the options for the endpoint (OPTIONS /api/environment)
func (h Handlers) OptionsEnvironment(ctx context.Context, r oapi.OptionsEnvironmentRequestObject) (res oapi.OptionsEnvironmentResponseObject, fault error) {
	return doOptionsEnvironment(ctx, r)
}

// CreateEnvironment Create a new environment (POST /api/environment)
func (h Handlers) CreateEnvironment(ctx context.Context, r oapi.CreateEnvironmentRequestObject) (res oapi.CreateEnvironmentResponseObject, fault error) {
	return doCreateEnvironment(ctx, r)
}

// DeleteEnvironment Delete an environment (DELETE /api/environment/{environment_name})
func (h Handlers) DeleteEnvironment(ctx context.Context, r oapi.DeleteEnvironmentRequestObject) (res oapi.DeleteEnvironmentResponseObject, fault error) {
	return doDeleteEnvironment(ctx, r)
}

// OptionsEnvironmentName Check the options for the endpoint (OPTIONS /api/environment/{environment_name})
func (h Handlers) OptionsEnvironmentName(ctx context.Context, r oapi.OptionsEnvironmentNameRequestObject) (res oapi.OptionsEnvironmentNameResponseObject, fault error) {
	return doOptionsEnvironmentName(ctx, r)
}

func (h Handlers) UICreateSingleFlag(ctx context.Context, r oapi.UICreateSingleFlagRequestObject) (res oapi.UICreateSingleFlagResponseObject, fault error) {
	return doUICreateSingleFlag(ctx, r)
}

func (h Handlers) UIOptionsCreateSingleFlag(ctx context.Context, r oapi.UIOptionsCreateSingleFlagRequestObject) (res oapi.UIOptionsCreateSingleFlagResponseObject, fault error) {
	return doUIOptionsCreateSingleFlag(ctx, r)
}

func (h Handlers) UIUpdateSingleFlag(ctx context.Context, r oapi.UIUpdateSingleFlagRequestObject) (res oapi.UIUpdateSingleFlagResponseObject, fault error) {
	return doUIUpdateSingleFlag(ctx, r)
}

func (h Handlers) UIOptionsSingleFlag(ctx context.Context, r oapi.UIOptionsSingleFlagRequestObject) (res oapi.UIOptionsSingleFlagResponseObject, fault error) {
	return doUIOptionsSingleFlag(ctx, r)
}

func (h Handlers) UIKillSingleFlag(ctx context.Context, r oapi.UIKillSingleFlagRequestObject) (res oapi.UIKillSingleFlagResponseObject, fault error) {
	return doUIKillSingleFlag(ctx, r)
}

// func (s Server) (ctx context.Context, r oapi.RequestObject) (res oapi.ResponseObject, fault error) {
// }

// func (s Server) (ctx context.Context, r oapi.RequestObject) (res oapi.ResponseObject, fault error) {
// }
