package cmd

import (
	"github.com/spf13/cobra"
)

var tableDescriptionsCmd = &cobra.Command{
	Use:   "table-descriptions",
	Short: "Manage Table Descriptions",
	Long:  `Commands for managing Table Descriptions resources.`,
}

func init() {
	datasourcesCmd.AddCommand(tableDescriptionsCmd)
}
