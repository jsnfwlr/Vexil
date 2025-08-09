package daemon

import (
	"github.com/spf13/cobra"
	otelTrace "go.opentelemetry.io/otel/trace"

	"github.com/jsnfwlr/o11y"
	"github.com/jsnfwlr/vexil/internal/api"
	"github.com/jsnfwlr/vexil/internal/db"
)

func init() {
	BaseCmd.AddCommand(StartCmd)
}

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "run the daemon",
	Run:   StartRun,
}

func StartRun(cmd *cobra.Command, args []string) {
	ctx, span := tracer.Start(cmd.Context(), "start", otelTrace.WithSpanKind(otelTrace.SpanKindClient))
	defer span.End()

	o := o11y.Get(ctx)
	o.Info("starting vexil daemon", span)

	dbConfig, err := db.LoadConfig()
	if err != nil {
		o.Fatal(err, span)
	}

	dbClient, err := db.Connect(ctx, dbConfig)
	if err != nil {
		o.Fatal(err, span)
	}

	defer func() {
		dbClient.Close()
	}()

	apiConfig, err := api.LoadConfig(dbClient)
	if err != nil {
		o.Fatal(err, span)
	}

	srvr, err := api.New(ctx, apiConfig)
	if err != nil {
		o.Fatal(err, span)
	}

	err = srvr.Start(ctx)
	if err != nil {
		o.Fatal(err, span)
	}
}
