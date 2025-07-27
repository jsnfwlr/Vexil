package curl

import "github.com/spf13/cobra"

var BaseCmd = &cobra.Command{
	Use:   "curl",
	Short: "make test calls to the API",
}
