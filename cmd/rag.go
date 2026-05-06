package cmd

import (
	"github.com/spf13/cobra"
)

var ragCmd = &cobra.Command{
	Use:   "rag",
	Short: "Manage rag",
	Long:  `Commands for managing rag resources.`,
}

func init() {
	datasourcesCmd.AddCommand(ragCmd)
}
