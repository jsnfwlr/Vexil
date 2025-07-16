package database

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	etc "github.com/jsnfwlr/vexil/etc/db/migrations"
	database "github.com/jsnfwlr/vexil/internal/db"
)

func init() {
	BaseCmd.AddCommand(MigrateCmd)
}

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "run migrations",
	Run:   MigrateRun,
}

func MigrateRun(cmd *cobra.Command, args []string) {
	ctx := cmd.Context()

	cfg, err := database.LoadConfig()
	if err != nil {
		fmt.Printf("failed to load config: %v\n", err)
		return
	}

	err = RunMigrate(ctx, cfg)
	if err != nil {
		fmt.Printf("Error running migrations: %v\n", err)
		return
	}
}

func RunMigrate(ctx context.Context, cfg database.ConfigProvider) (fault error) {
	m, err := database.NewMigrator(ctx, cfg, etc.Migrations)
	if err != nil {
		return fmt.Errorf("failed to create migrator: %w", err)
	}

	if err := m.Migrate(); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}
