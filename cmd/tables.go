package cmd

import (
	"github.com/spf13/cobra"
)

var tablesCmd = &cobra.Command{
	Use:   "tables",
	Short: "Manage Tables",
	Long:  `Commands for managing Tables resources.`,
}

func init() {
	datasourcesCmd.AddCommand(tablesCmd)
}
