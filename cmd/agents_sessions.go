package cmd

import (
	"github.com/spf13/cobra"
)

var agentsSessionsCmd = &cobra.Command{
	Use:   "agents-sessions",
	Short: "Manage AgentsSessions",
	Long:  `Commands for managing AgentsSessions resources.`,
}

func init() {
	agentsCmd.AddCommand(agentsSessionsCmd)
}
