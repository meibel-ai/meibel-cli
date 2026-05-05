package cmd

import (
	"github.com/spf13/cobra"
)

var documentsCmd = &cobra.Command{
	Use:   "documents",
	Short: "Manage Documents",
	Long:  `Commands for managing Documents resources.`,
}

func init() {
	rootCmd.AddCommand(documentsCmd)
}
