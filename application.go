package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/takaaki-mizuno/go-gin-template/cmd/app"
	"github.com/takaaki-mizuno/go-gin-template/cmd/config"
	"github.com/takaaki-mizuno/go-gin-template/cmd/db"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "CLI for Go Gin Template",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "App server related",
}

var appServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run application server",
	RunE:  app.Serve,
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

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Config related commands",
}

var configPrintCmd = &cobra.Command{
	Use:   "print",
	Short: "Output all config values",
	RunE:  config.Print,
}

func init() {
	rootCmd.AddCommand(appCmd)
	appCmd.AddCommand(appServeCmd)

	rootCmd.AddCommand(dbCmd)
	dbCmd.AddCommand(dbMigrateCmd)
	dbCmd.AddCommand(dbSeedCmd)

	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configPrintCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
