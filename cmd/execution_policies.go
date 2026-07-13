package cmd

import (
	"github.com/spf13/cobra"
)

var executionPoliciesCmd = &cobra.Command{
	Use:   "execution-policies",
	Short: "Manage ExecutionPolicies",
	Long:  `Commands for managing ExecutionPolicies resources.`,
}

func init() {
	rootCmd.AddCommand(executionPoliciesCmd)
}
