package cmd

import (
	"github.com/spf13/cobra"
)

var batchDefinitionsCmd = &cobra.Command{
	Use:   "batch-definitions",
	Short: "Manage BatchDefinitions",
	Long:  `Commands for managing BatchDefinitions resources.`,
}

func init() {
	rootCmd.AddCommand(batchDefinitionsCmd)
}
