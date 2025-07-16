package event

import "github.com/spf13/cobra"

var BaseCmd = &cobra.Command{
	Use:   "event",
	Short: "trigger an event for the SSE endpoint",
}
