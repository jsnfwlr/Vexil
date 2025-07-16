package handlers

import (
	"context"

	"github.com/jsnfwlr/vexil/internal/api/oapi"
)

// BulkFlags Get bulk list of flags (GET /api/flag)
func (s Server) BulkFlags(ctx context.Context, r oapi.BulkFlagsRequestObject) (res oapi.BulkFlagsResponseObject, fault error) {
	return doBulkFlags(ctx, r)
}

func doBulkFlags(ctx context.Context, r oapi.BulkFlagsRequestObject) (res oapi.BulkFlagsResponseObject, fault error) {
	return oapi.BulkFlags200JSONResponse{}, nil
}

// OptionsFlag Check the options for the endpoint (OPTIONS /api/flag)
func (s Server) OptionsFlag(ctx context.Context, r oapi.OptionsFlagRequestObject) (res oapi.OptionsFlagResponseObject, fault error) {
	return doOptionsFlag(ctx, r)
}

func doOptionsFlag(ctx context.Context, r oapi.OptionsFlagRequestObject) (res oapi.OptionsFlagResponseObject, fault error) {
	return oapi.OptionsFlag200Response{}, nil
}

// CreateFlag Create a new flag (POST /api/flag)
func (s Server) CreateFlag(ctx context.Context, r oapi.CreateFlagRequestObject) (res oapi.CreateFlagResponseObject, fault error) {
	return doCreateFlag(ctx, r)
}

func doCreateFlag(ctx context.Context, r oapi.CreateFlagRequestObject) (res oapi.CreateFlagResponseObject, fault error) {
	return oapi.CreateFlag201JSONResponse{}, nil
}

// DeleteFlag Delete a flag (DELETE /api/flag/{flag_id})
func (s Server) DeleteFlag(ctx context.Context, r oapi.DeleteFlagRequestObject) (res oapi.DeleteFlagResponseObject, fault error) {
	return doDeleteFlag(ctx, r)
}

func doDeleteFlag(ctx context.Context, r oapi.DeleteFlagRequestObject) (res oapi.DeleteFlagResponseObject, fault error) {
	return oapi.DeleteFlag204Response{}, nil
}

// OptionsFlagId Check the options for the endpoint (OPTIONS /api/flag/{flag_id})
func (s Server) OptionsFlagId(ctx context.Context, r oapi.OptionsFlagIdRequestObject) (res oapi.OptionsFlagIdResponseObject, fault error) {
	return doOptionsFlagId(ctx, r)
}

func doOptionsFlagId(ctx context.Context, r oapi.OptionsFlagIdRequestObject) (res oapi.OptionsFlagIdResponseObject, fault error) {
	return oapi.OptionsFlagId200Response{}, nil
}

// UpdateFlag Update a flag (PUT /api/flag/{flag_id})
func (s Server) UpdateFlag(ctx context.Context, r oapi.UpdateFlagRequestObject) (res oapi.UpdateFlagResponseObject, fault error) {
	return doUpdateFlag(ctx, r)
}

func doUpdateFlag(ctx context.Context, r oapi.UpdateFlagRequestObject) (res oapi.UpdateFlagResponseObject, fault error) {
	return oapi.UpdateFlag200JSONResponse{}, nil
}

// OptionsEnvironmentNameFlagId Check the options for the endpoint (OPTIONS /api/environment/{environment_name}/flag/{flag_id})
func (s Server) OptionsEnvironmentNameFlagId(ctx context.Context, r oapi.OptionsEnvironmentNameFlagIdRequestObject) (res oapi.OptionsEnvironmentNameFlagIdResponseObject, fault error) {
	return doOptionsEnvironmentNameFlagId(ctx, r)
}

func doOptionsEnvironmentNameFlagId(ctx context.Context, r oapi.OptionsEnvironmentNameFlagIdRequestObject) (res oapi.OptionsEnvironmentNameFlagIdResponseObject, fault error) {
	return oapi.OptionsEnvironmentNameFlagId200Response{}, nil
}

// UpdateEnvironmentFlagValue Update a flag in an environment (PUT /api/environment/{environment_name}/flag/{flag_id})
func (s Server) UpdateEnvironmentFlagValue(ctx context.Context, r oapi.UpdateEnvironmentFlagValueRequestObject) (res oapi.UpdateEnvironmentFlagValueResponseObject, fault error) {
	return doUpdateEnvironmentFlagValue(ctx, r)
}

func doUpdateEnvironmentFlagValue(ctx context.Context, r oapi.UpdateEnvironmentFlagValueRequestObject) (res oapi.UpdateEnvironmentFlagValueResponseObject, fault error) {
	return oapi.UpdateEnvironmentFlagValue200JSONResponse{}, nil
}
