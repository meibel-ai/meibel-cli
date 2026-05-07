package cmd

import (
	"github.com/spf13/cobra"
)

var sessionsCmd = &cobra.Command{
	Use:   "sessions",
	Short: "Manage Sessions",
	Long:  `Commands for managing Sessions resources.`,
}

func init() {
	agentsCmd.AddCommand(sessionsCmd)
}
