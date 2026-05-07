package cmd

import (
	"github.com/spf13/cobra"
)

var promptsCmd = &cobra.Command{
	Use:   "prompts",
	Short: "Manage Prompts",
	Long:  `Commands for managing Prompts resources.`,
}

func init() {
	rootCmd.AddCommand(promptsCmd)
}
