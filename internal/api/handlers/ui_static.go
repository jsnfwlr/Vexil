package handlers

import (
	"context"

	"github.com/jsnfwlr/vexil/internal/api/oapi"
)

// GetUiIndex Get the index.html for the UI (GET /ui)
func (s Server) GetUiIndex(ctx context.Context, r oapi.GetUiIndexRequestObject) (res oapi.GetUiIndexResponseObject, fault error) {
	return doGetUiIndex(ctx, r)
}

func doGetUiIndex(ctx context.Context, r oapi.GetUiIndexRequestObject) (res oapi.GetUiIndexResponseObject, fault error) {
	return oapi.GetUiIndex200TexthtmlResponse{}, nil
}

// OptionsUi Check the options for the endpoint (OPTIONS /ui)
func (s Server) OptionsUi(ctx context.Context, r oapi.OptionsUiRequestObject) (res oapi.OptionsUiResponseObject, fault error) {
	return doOptionsUi(ctx, r)
}

func doOptionsUi(ctx context.Context, r oapi.OptionsUiRequestObject) (res oapi.OptionsUiResponseObject, fault error) {
	return oapi.OptionsUi200Response{}, nil
}

// GetUiFile Get the static files for the UI (GET /ui/{file})
func (s Server) GetUiFile(ctx context.Context, r oapi.GetUiFileRequestObject) (res oapi.GetUiFileResponseObject, fault error) {
	return doGetUiFile(ctx, r)
}

func doGetUiFile(ctx context.Context, r oapi.GetUiFileRequestObject) (res oapi.GetUiFileResponseObject, fault error) {
	return oapi.GetUiFile200TexthtmlResponse{}, nil
}

// OptionsUiFile Check the options for the endpoint (OPTIONS /ui/{file})
func (s Server) OptionsUiFile(ctx context.Context, r oapi.OptionsUiFileRequestObject) (res oapi.OptionsUiFileResponseObject, fault error) {
	return doOptionsUiFile(ctx, r)
}

func doOptionsUiFile(ctx context.Context, r oapi.OptionsUiFileRequestObject) (res oapi.OptionsUiFileResponseObject, fault error) {
	return oapi.OptionsUiFile200Response{}, nil
}
