package cmd

import (
	"github.com/spf13/cobra"
)

var datasourcesCmd = &cobra.Command{
	Use:   "datasources",
	Short: "Manage Datasources",
	Long:  `Commands for managing Datasources resources.`,
}

func init() {
	rootCmd.AddCommand(datasourcesCmd)
}
