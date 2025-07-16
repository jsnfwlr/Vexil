package cmd

import (
	"context"
	"os"

	"github.com/jsnfwlr/vexil/internal/cmd/daemon"
	"github.com/jsnfwlr/vexil/internal/cmd/database"
	"github.com/jsnfwlr/vexil/internal/cmd/event"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Version: "v0.1.0",
	Use:     "vexil",
	Short:   "vexil CLI",
}

func Execute(ctx context.Context) {
	RootCmd.SetContext(ctx)
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	_ = godotenv.Load("dev.env", ".env")

	cobra.EnableCaseInsensitive = true
	// cobra.EnableCommandSorting = false

	RootCmd.CompletionOptions.DisableDefaultCmd = true
	RootCmd.Flags().SortFlags = false
	RootCmd.PersistentFlags().SortFlags = false

	RootCmd.AddCommand(daemon.BaseCmd)
	RootCmd.AddCommand(database.BaseCmd)
	RootCmd.AddCommand(event.BaseCmd)
}
