package cmd

import (
	"github.com/spf13/cobra"
)

var datasourcesCmd = &cobra.Command{
	Use:   "datasources",
	Short: "Manage datasources",
	Long:  `Commands for managing datasources resources.`,
}

func init() {
	rootCmd.AddCommand(datasourcesCmd)
}
