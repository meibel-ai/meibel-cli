package cmd

import (
	"github.com/spf13/cobra"
)

var executionsCmd = &cobra.Command{
	Use:   "executions",
	Short: "Manage executions",
	Long:  `Commands for managing executions resources.`,
}

func init() {
	blueprintsCmd.AddCommand(executionsCmd)
}
