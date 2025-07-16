package handlers

import (
	"context"

	"github.com/jsnfwlr/vexil/internal/api/oapi"
)

// GetFlagsByEnvironment Get flags by environment (GET /api/environment/{environment_name}/flag)
func (s Server) GetFlagsByEnvironment(ctx context.Context, r oapi.GetFlagsByEnvironmentRequestObject) (res oapi.GetFlagsByEnvironmentResponseObject, fault error) {
	return doGetFlagsByEnvironment(ctx, r)
}

func doGetFlagsByEnvironment(ctx context.Context, r oapi.GetFlagsByEnvironmentRequestObject) (res oapi.GetFlagsByEnvironmentResponseObject, fault error) {
	var out oapi.GetFlagsByEnvironment200JSONResponse

	out = append(out, oapi.EnvironmentFlag{
		Name:  "FIRST_STRING",
		Type:  "string",
		Value: "13cm of thirty stand green wool",
	})
	out = append(out, oapi.EnvironmentFlag{
		Name:  "SECOND_STRING",
		Type:  "string",
		Value: "30cm of three stand blue cotton",
	})

	return out, nil
}

// OptionsEnvironmentNameFlag Check the options for the endpoint (OPTIONS /api/environment/{environment_name}/flag)
func (s Server) OptionsEnvironmentNameFlag(ctx context.Context, r oapi.OptionsEnvironmentNameFlagRequestObject) (res oapi.OptionsEnvironmentNameFlagResponseObject, fault error) {
	return doOptionsEnvironmentNameFlag(ctx, r)
}

func doOptionsEnvironmentNameFlag(ctx context.Context, r oapi.OptionsEnvironmentNameFlagRequestObject) (res oapi.OptionsEnvironmentNameFlagResponseObject, fault error) {
	return oapi.OptionsEnvironmentNameFlag200Response{}, nil
}
