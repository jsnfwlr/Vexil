package db_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	etc "github.com/jsnfwlr/vexil/etc/db/migrations"
	migrations "github.com/jsnfwlr/vexil/internal/db"
)

func TestFileSystem(t *testing.T) {
	fs := etc.Migrations
	em := migrations.MigrationFS{FS: fs}

	fi, err := em.ReadDir(".")
	if err != nil {
		t.Fatalf("could not read the directory: %v", err)
	}

	if len(fi) == 0 {
		t.Fatalf("no files found in the directory")
	}

	for _, f := range fi {
		t.Logf("name: %s, size: %d, mode: %s, modTime: %v, isDir: %t", f.Name(), f.Size(), f.Mode(), f.ModTime(), f.IsDir())
	}

	sharedPaths, err := em.Glob(filepath.Join("*", "*.sql"))
	if err != nil {
		t.Errorf("could not get globs: %s", err)
	}

	for _, p := range sharedPaths {
		t.Logf("path: %s", p)
	}
}

func TestMigrator(t *testing.T) {
	fs := etc.Migrations
	ctx := context.Background()

	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_DATABASE", "vexil")
	os.Setenv("POSTGRES_USER", "vexil")
	os.Setenv("POSTGRES_PASSWORD", "vexil")

	cfg, err := migrations.LoadConfig()
	if err != nil {
		t.Fatalf("could not load the configuration: %v", err)
	}

	m, err := migrations.NewMigrator(ctx, cfg, fs)
	if err != nil {
		t.Fatalf("could not create the migrator: %v", err)
	}

	i, err := m.Info(-1)
	if err != nil {
		t.Fatalf("could not get the migration info: %v", err)
	}

	t.Logf("host: %s:%s, database: %s, currentVersion: %d, targetVersion: %d\n%s", i.Hostname, i.Port, i.Database, i.Migrations.CurrentVersion, i.Migrations.TargetVersion, i.Migrations.Summary)

	err = migrations.RunMigrations(ctx, cfg, fs, -1, true)
	if err != nil {
		t.Fatalf("could not run the migrations: %v", err)
	}
}
