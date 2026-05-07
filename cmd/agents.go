package cmd

import (
	"github.com/spf13/cobra"
)

var agentsCmd = &cobra.Command{
	Use:   "agents",
	Short: "Manage Agents",
	Long:  `Commands for managing Agents resources.`,
}

func init() {
	rootCmd.AddCommand(agentsCmd)
}
