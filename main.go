package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "CLI for Redimir",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var appCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run application server",
	RunE:  app.NewAppServer,
}

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Manipulate the database",
}

var dbMigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run the migrations",
	RunE:  db.Migrate,
}

var dbSeedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Run the migrations",
	RunE:  db.Seed,
}

func init() {
	rootCmd.AddCommand(dbCmd)
	rootCmd.AddCommand(appCmd)
	dbCmd.AddCommand(dbMigrateCmd)
	dbCmd.AddCommand(dbSeedCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
