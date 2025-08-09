package daemon

import (
	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("github.com/jsnfwlr/vexil/cmd/daemon")

var BaseCmd = &cobra.Command{
	Use:   "daemon",
	Short: "start vexil as a daemon",
}
