package handlers

import (
	"context"

	"github.com/jsnfwlr/vexil/internal/api/oapi"
)

// ListEnvironments List environments (GET /api/environment)
func (s Server) ListEnvironments(ctx context.Context, r oapi.ListEnvironmentsRequestObject) (res oapi.ListEnvironmentsResponseObject, fault error) {
	return doListEnvironments(ctx, r)
}

func doListEnvironments(ctx context.Context, r oapi.ListEnvironmentsRequestObject) (res oapi.ListEnvironmentsResponseObject, fault error) {
	return oapi.ListEnvironments200JSONResponse{}, nil
}

// OptionsEnvironment Check the options for the endpoint (OPTIONS /api/environment)
func (s Server) OptionsEnvironment(ctx context.Context, r oapi.OptionsEnvironmentRequestObject) (res oapi.OptionsEnvironmentResponseObject, fault error) {
	return doOptionsEnvironment(ctx, r)
}

func doOptionsEnvironment(ctx context.Context, r oapi.OptionsEnvironmentRequestObject) (res oapi.OptionsEnvironmentResponseObject, fault error) {
	return oapi.OptionsEnvironment200Response{}, nil
}

// CreateEnvironment Create a new environment (POST /api/environment)
func (s Server) CreateEnvironment(ctx context.Context, r oapi.CreateEnvironmentRequestObject) (res oapi.CreateEnvironmentResponseObject, fault error) {
	return doCreateEnvironment(ctx, r)
}

func doCreateEnvironment(ctx context.Context, r oapi.CreateEnvironmentRequestObject) (res oapi.CreateEnvironmentResponseObject, fault error) {
	return oapi.CreateEnvironment201JSONResponse{}, nil
}

// DeleteEnvironment Delete an environment (DELETE /api/environment/{environment_name})
func (s Server) DeleteEnvironment(ctx context.Context, r oapi.DeleteEnvironmentRequestObject) (res oapi.DeleteEnvironmentResponseObject, fault error) {
	return doDeleteEnvironment(ctx, r)
}

func doDeleteEnvironment(ctx context.Context, r oapi.DeleteEnvironmentRequestObject) (res oapi.DeleteEnvironmentResponseObject, fault error) {
	return oapi.DeleteEnvironment204Response{}, nil
}

// OptionsEnvironmentName Check the options for the endpoint (OPTIONS /api/environment/{environment_name})
func (s Server) OptionsEnvironmentName(ctx context.Context, r oapi.OptionsEnvironmentNameRequestObject) (res oapi.OptionsEnvironmentNameResponseObject, fault error) {
	return doOptionsEnvironmentName(ctx, r)
}

func doOptionsEnvironmentName(ctx context.Context, r oapi.OptionsEnvironmentNameRequestObject) (res oapi.OptionsEnvironmentNameResponseObject, fault error) {
	return oapi.OptionsEnvironmentName200Response{}, nil
}
