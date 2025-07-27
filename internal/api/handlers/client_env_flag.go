package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jsnfwlr/o11y"
	"github.com/jsnfwlr/vexil/internal/api/oapi"
	"github.com/jsnfwlr/vexil/internal/db"
	"github.com/jsnfwlr/vexil/internal/log"
)

type FlagsByEnvQueryProvider interface {
	GetFlagsByEnvironmentName(ctx context.Context, name string) ([]db.GetFlagsByEnvironmentNameRow, error)
}

func doGetFlagsByEnvironment(ctx context.Context, q FlagsByEnvQueryProvider, r oapi.GetFlagsByEnvironmentRequestObject) (res oapi.GetFlagsByEnvironmentResponseObject, fault error) {
	var out oapi.GetFlagsByEnvironment200JSONResponse

	ctx, o := o11y.Get(ctx, nil)

	rows, err := q.GetFlagsByEnvironmentName(ctx, r.EnvironmentName)
	if err != nil {
		o.Error(err, log.EnvironmentNameKey, r.EnvironmentName)
		return oapi.GetFlagsByEnvironment400JSONResponse{
			Error:     err.Error(),
			ErrorCode: http.StatusBadRequest,
		}, err
	}

	if len(rows) == 0 {
		err = fmt.Errorf("no flags found for %s", r.EnvironmentName)
		return oapi.GetFlagsByEnvironment400JSONResponse{
			Error:     err.Error(),
			ErrorCode: http.StatusBadRequest,
		}, err
	}

	for _, f := range rows {
		vt, err := f.FeatureFlag.ValueType.ToAPIEnum()
		if err != nil {
			o.Error(err, log.EnvironmentNameKey, r.EnvironmentName)
			return oapi.GetFlagsByEnvironment500JSONResponse{
				Error:     err.Error(),
				ErrorCode: http.StatusInternalServerError,
			}, err
		}
		of := oapi.EnvironmentFlag{
			Name:  f.FeatureFlag.Name,
			Type:  vt,
			Value: f.FeatureFlagValue.Value,
		}

		out = append(out, of)
	}

	return out, nil
}

func doOptionsEnvironmentNameFlag(ctx context.Context, r oapi.OptionsEnvironmentNameFlagRequestObject) (res oapi.OptionsEnvironmentNameFlagResponseObject, fault error) {
	return oapi.OptionsEnvironmentNameFlag200Response{}, nil
}
