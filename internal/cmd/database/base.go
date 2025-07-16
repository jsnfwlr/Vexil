package database

import "github.com/spf13/cobra"

var BaseCmd = &cobra.Command{
	Use:   "database",
	Short: "manage vexil's database",
}
