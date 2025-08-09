package daemon

import (
	"errors"
	"net/http"

	"github.com/spf13/cobra"
	otelTrace "go.opentelemetry.io/otel/trace"

	"github.com/jsnfwlr/vexil/internal/log"

	"github.com/jsnfwlr/o11y"
)

func init() {
	BaseCmd.AddCommand(CheckCmd)
}

var CheckCmd = &cobra.Command{
	Use:   "check",
	Short: "check the daemon is running",
	Run:   CheckRun,
}

func CheckRun(cmd *cobra.Command, args []string) {
	ctx, span := tracer.Start(cmd.Context(), "healthcheck", otelTrace.WithSpanKind(otelTrace.SpanKindClient))
	defer span.End()

	o := o11y.Get(ctx)
	o.Debug("checking vexil daemon", span)

	client := &http.Client{
		Transport: http.DefaultTransport,
	}
	err := o11y.AddLoggingToHTTPClient(client)
	if err != nil {
		o.Error(err, span)
		return
	}

	req, err := http.NewRequest(http.MethodGet, "http://localhost:9765/ping", nil)
	if err != nil {
		o.Error(err, span)
		return
	}

	req.Header.Set("x-span-id", o.SpanContext().SpanID().String())
	req.Header.Set("x-trace-id", o.SpanContext().TraceID().String())

	resp, err := client.Do(req)
	if err != nil {
		o.Error(err, span)
	}

	if resp.StatusCode != http.StatusOK {
		o.Error(errors.New("healthcheck encountered invalid status code"), span, log.StatusCodeKey, resp.StatusCode)
	}

	o.Debug("vexil daemon is running", span, log.StatusCodeKey, resp.StatusCode)
}
