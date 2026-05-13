package cmd

import (
	"github.com/spf13/cobra"
)

var batchesCmd = &cobra.Command{
	Use:   "batches",
	Short: "Manage Batches",
	Long:  `Commands for managing Batches resources.`,
}

func init() {
	rootCmd.AddCommand(batchesCmd)
}
