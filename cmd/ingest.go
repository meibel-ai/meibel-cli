package cmd

import (
	"github.com/spf13/cobra"
)

var ingestCmd = &cobra.Command{
	Use:   "ingest",
	Short: "Manage Ingest",
	Long:  `Commands for managing Ingest resources.`,
}

func init() {
	datasourcesCmd.AddCommand(ingestCmd)
}
