package daemon

import (
	"errors"
	"net/http"

	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel"
	otelTrace "go.opentelemetry.io/otel/trace"

	"github.com/jsnfwlr/vexil/internal/log"

	"github.com/jsnfwlr/o11y"
)

var tracer = otel.Tracer("github.com/jsnfwlr/vexil/cmd/daemon")

func init() {
	BaseCmd.AddCommand(CheckCmd)
}

var CheckCmd = &cobra.Command{
	Use:   "check",
	Short: "check the daemon is running",
	Run:   CheckRun,
}

func CheckRun(cmd *cobra.Command, args []string) {
	_, o := o11y.Get(tracer.Start(cmd.Context(), "healthcheck", otelTrace.WithSpanKind(otelTrace.SpanKindClient)))

	defer o.End()

	o.Debug("checking vexil daemon")

	client := &http.Client{
		Transport: http.DefaultTransport,
	}
	err := o11y.AddLoggingToHTTPClient(client)
	if err != nil {
		o.Error(err)
		return
	}

	req, err := http.NewRequest(http.MethodGet, "http://localhost:9765/ping", nil)
	if err != nil {
		o.Error(err)
		return
	}

	req.Header.Set("x-span-id", o.SpanContext().SpanID().String())
	req.Header.Set("x-trace-id", o.SpanContext().TraceID().String())

	resp, err := client.Do(req)
	if err != nil {
		o.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		o.Error(errors.New("healthcheck encountered invalid status code"), log.StatusCodeKey, resp.StatusCode)
	}

	o.Debug("vexil daemon is running", log.StatusCodeKey, resp.StatusCode)
}
