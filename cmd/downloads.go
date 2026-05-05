package cmd

import (
	"github.com/spf13/cobra"
)

var downloadsCmd = &cobra.Command{
	Use:   "downloads",
	Short: "Manage Downloads",
	Long:  `Commands for managing Downloads resources.`,
}

func init() {
	rootCmd.AddCommand(downloadsCmd)
}
