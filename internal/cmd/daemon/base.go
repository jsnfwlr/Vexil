package daemon

import "github.com/spf13/cobra"

var BaseCmd = &cobra.Command{
	Use:   "daemon",
	Short: "start vexil as a daemon",
}
