package cmd

import (
	"github.com/spf13/cobra"
)

var executionsCmd = &cobra.Command{
	Use:   "executions",
	Short: "Manage Executions",
	Long:  `Commands for managing Executions resources.`,
}

func init() {
	batchesCmd.AddCommand(executionsCmd)
}
