package cmd

import (
	"github.com/spf13/cobra"
)

var contentCmd = &cobra.Command{
	Use:   "content",
	Short: "Manage Content",
	Long:  `Commands for managing Content resources.`,
}

func init() {
	rootCmd.AddCommand(contentCmd)
}
