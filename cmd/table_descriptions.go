package cmd

import (
	"github.com/spf13/cobra"
)

var tableDescriptionsCmd = &cobra.Command{
	Use:   "table-descriptions",
	Short: "Manage TableDescriptions",
	Long:  `Commands for managing TableDescriptions resources.`,
}

func init() {
	rootCmd.AddCommand(tableDescriptionsCmd)
}
