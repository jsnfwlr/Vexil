package handlers

import (
	"context"

	"github.com/jsnfwlr/vexil/internal/api/oapi"
)

func doListEnvironments(ctx context.Context, r oapi.ListEnvironmentsRequestObject) (res oapi.ListEnvironmentsResponseObject, fault error) {
	return oapi.ListEnvironments200JSONResponse{}, nil
}

func doOptionsEnvironment(ctx context.Context, r oapi.OptionsEnvironmentRequestObject) (res oapi.OptionsEnvironmentResponseObject, fault error) {
	return oapi.OptionsEnvironment200Response{}, nil
}

func doCreateEnvironment(ctx context.Context, r oapi.CreateEnvironmentRequestObject) (res oapi.CreateEnvironmentResponseObject, fault error) {
	return oapi.CreateEnvironment201JSONResponse{}, nil
}

func doDeleteEnvironment(ctx context.Context, r oapi.DeleteEnvironmentRequestObject) (res oapi.DeleteEnvironmentResponseObject, fault error) {
	return oapi.DeleteEnvironment204Response{}, nil
}

func doOptionsEnvironmentName(ctx context.Context, r oapi.OptionsEnvironmentNameRequestObject) (res oapi.OptionsEnvironmentNameResponseObject, fault error) {
	return oapi.OptionsEnvironmentName200Response{}, nil
}
