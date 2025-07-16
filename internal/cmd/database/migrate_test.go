package database_test

import (
	"context"
	"testing"

	cmd "github.com/jsnfwlr/vexil/internal/cmd/database"
	database "github.com/jsnfwlr/vexil/internal/db"
)

func TestMigrateRun(t *testing.T) {
	ctx := context.Background()

	cfg := database.Config{
		Host:     "localhost",
		Port:     "5432",
		Database: "vexil",
		Username: "vexil",
		Password: "vexil",
	}

	err := cmd.RunMigrate(ctx, cfg)
	if err != nil {
		t.Fatalf("migrateRun failed: %v", err)
	}

	// Additional checks can be added here to verify the state of the database
	// after migration, such as checking for specific tables or data.
}
