package daemon

import (
	"github.com/spf13/cobra"

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
	ctx, o := o11y.Get(cmd.Context(), nil)

	o.Info("starting vexil daemon")

	dbConfig, err := db.LoadConfig()
	if err != nil {
		o.Fatal(err)
	}

	dbClient, err := db.Connect(ctx, dbConfig)
	if err != nil {
		o.Fatal(err)
	}

	defer func() {
		dbClient.Close()
	}()

	apiConfig, err := api.LoadConfig(dbClient)
	if err != nil {
		o.Fatal(err)
	}

	srvr, err := api.New(ctx, apiConfig)
	if err != nil {
		o.Fatal(err)
	}

	err = srvr.Start(ctx)
	if err != nil {
		o.Fatal(err)
	}
}
