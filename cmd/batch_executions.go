package cmd

import (
	"github.com/spf13/cobra"
)

var batchExecutionsCmd = &cobra.Command{
	Use:   "batch-executions",
	Short: "Manage BatchExecutions",
	Long:  `Commands for managing BatchExecutions resources.`,
}

func init() {
	rootCmd.AddCommand(batchExecutionsCmd)
}
