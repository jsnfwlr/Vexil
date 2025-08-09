package handlers

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jsnfwlr/o11y"
	"github.com/jsnfwlr/vexil/internal/api/oapi"
	"github.com/jsnfwlr/vexil/internal/db"
	"github.com/jsnfwlr/vexil/internal/templates"
	otelTrace "go.opentelemetry.io/otel/trace"
)

type Table struct {
	Pages        Pagination
	Environments []string
	Flags        []EnvFlag
}

type Pagination struct {
	Sort      string
	Direction string
	Size      int
	Total     int
	Rows      int
	Page      int
}

type EnvFlag struct {
	Id      int32
	Name    string
	Type    string
	Default string
	Values  Flag
}

type FindBulkFlagsQueryProvider interface {
	PageFlags(ctx context.Context, arg db.PageFlagsParams) ([]db.PageFlagsRow, error)
}

func doUIFindBulkFlags(ctx context.Context, dbClient FindBulkFlagsQueryProvider, r oapi.UIFindBulkFlagsRequestObject) (res oapi.UIFindBulkFlagsResponseObject, fault error) {
	ctx, span := tracer.Start(ctx, "doUIFindBulkFlags", otelTrace.WithSpanKind(otelTrace.SpanKindServer))
	defer span.End()

	o := o11y.Get(ctx)
	o.Debug("getting flag table", span)

	var b []byte
	var op io.Reader

	arg := db.PageFlagsParams{
		Page: 0,
		Size: 25,
		Sort: pgtype.Text{String: "name asc", Valid: true},
	}
	pages := Pagination{
		Page:      0,
		Size:      25,
		Sort:      "name",
		Direction: "asc",
		Total:     0,
		Rows:      0,
	}

	if r.Params.Page != nil {
		arg.Page = int64(*r.Params.Page)
		pages.Page = *r.Params.Page
	}
	if r.Params.Size != nil {
		arg.Size = int64(*r.Params.Size)
		pages.Size = *r.Params.Size
	}

	if r.Params.SortBy != nil {
		dir := "asc"
		if r.Params.SortDirection != nil {
			dir = string(*r.Params.SortDirection)
		}

		arg.Sort = pgtype.Text{String: fmt.Sprintf("%s %s", string(*r.Params.SortBy), dir), Valid: true}
		pages.Sort = string(*r.Params.SortBy)
		pages.Direction = dir

	}

	flags := []EnvFlag{}
	results, err := dbClient.PageFlags(ctx, arg)
	if err != nil {
		b = []byte(`<div id="error">Could not list flags</div>`)
		op = bytes.NewReader(b)
		return oapi.UIFindBulkFlags500TexthtmlResponse{Body: op, ContentLength: int64(len(b))}, nil
	}

	for _, res := range results {
		// f := EnvFlag{
		// 	Id: res.FeatureFlag.ID,
		// 	Name: res.FeatureFlag.Name,
		// 	Type: string(res.FeatureFlag.ValueType),
		// 	Default: res.FeatureFlag.DefaultValue,
		// }

		fmt.Printf("Name: %s Id: %d - Env: %s Val: %s", res.FeatureFlag.Name, res.FeatureFlag.ID, res.Environment.Name, res.FeatureFlagValue.Value)
	}

	data := Table{
		Pages: pages,
		Flags: flags,
	}

	funcs := template.FuncMap{
		"Loop": func(count int) []int {
			var i int
			var Items []int
			for i = 0; i < count; i++ {
				Items = append(Items, i)
			}
			return Items
		},
		"Add": func(val1, val2 int) int { return val1 + val2 },
	}

	tmpl, err := template.New("table.html").Funcs(funcs).ParseFS(templates.Files, "table.html")
	if err != nil {
		b = []byte(fmt.Sprintf(`<div id="error">Could not read template: %s</div>`, err))
		o.Error(err, span)
		op = bytes.NewReader(b)
		return oapi.UIFindBulkFlags500TexthtmlResponse{Body: op, ContentLength: int64(len(b))}, nil
	}

	var buf bytes.Buffer
	ip := bufio.NewWriter(&buf)
	err = tmpl.Execute(ip, data)
	if err != nil {
		b = []byte(fmt.Sprintf(`<div id="error">Could not render template: %s</div>`, err))
		o.Error(err, span)
		op = bytes.NewReader(b)
		return oapi.UIFindBulkFlags500TexthtmlResponse{Body: op, ContentLength: int64(len(b))}, nil
	}

	b = []byte(buf.Bytes())
	op = bytes.NewReader(b)

	// t := Table{}

	return oapi.UIFindBulkFlags200TexthtmlResponse{Body: op, ContentLength: int64(len(b))}, nil
}

func doUICreateSingleFlag(ctx context.Context, r oapi.UICreateSingleFlagRequestObject) (res oapi.UICreateSingleFlagResponseObject, fault error) {
	ctx, span := tracer.Start(ctx, "doUICreateSingleFlag", otelTrace.WithSpanKind(otelTrace.SpanKindServer))
	defer span.End()

	return oapi.UICreateSingleFlag200TexthtmlResponse{}, nil
}

func doUIUpdateSingleFlag(ctx context.Context, r oapi.UIUpdateSingleFlagRequestObject) (res oapi.UIUpdateSingleFlagResponseObject, fault error) {
	ctx, span := tracer.Start(ctx, "doUIUpdateSingleFlag", otelTrace.WithSpanKind(otelTrace.SpanKindServer))
	defer span.End()

	return oapi.UIUpdateSingleFlag200TexthtmlResponse{}, nil
}

func doUIKillSingleFlag(ctx context.Context, r oapi.UIKillSingleFlagRequestObject) (res oapi.UIKillSingleFlagResponseObject, fault error) {
	ctx, span := tracer.Start(ctx, "doUIKillSingleFlag", otelTrace.WithSpanKind(otelTrace.SpanKindServer))
	defer span.End()

	return oapi.UIKillSingleFlag200TexthtmlResponse{}, nil
}

// OPTIONS

func doUIOptionsFindFlags(ctx context.Context, r oapi.UIOptionsFindBulkFlagsRequestObject) (res oapi.UIOptionsFindBulkFlagsResponseObject, fault error) {
	ctx, span := tracer.Start(ctx, "doUIOptionsFindFlags", otelTrace.WithSpanKind(otelTrace.SpanKindServer))
	defer span.End()

	return oapi.UIOptionsFindBulkFlags200Response{}, nil
}

func doUIOptionsCreateSingleFlag(ctx context.Context, r oapi.UIOptionsCreateSingleFlagRequestObject) (res oapi.UIOptionsCreateSingleFlagResponseObject, fault error) {
	ctx, span := tracer.Start(ctx, "doUIOptionsCreateSingleFlag", otelTrace.WithSpanKind(otelTrace.SpanKindServer))
	defer span.End()

	return oapi.UIOptionsCreateSingleFlag200Response{}, nil
}

func doUIOptionsSingleFlag(ctx context.Context, r oapi.UIOptionsSingleFlagRequestObject) (res oapi.UIOptionsSingleFlagResponseObject, fault error) {
	ctx, span := tracer.Start(ctx, "doUIOptionsSingleFlag", otelTrace.WithSpanKind(otelTrace.SpanKindServer))
	defer span.End()

	return oapi.UIOptionsSingleFlag200Response{}, nil
}
