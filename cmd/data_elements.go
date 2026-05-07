package cmd

import (
	"github.com/spf13/cobra"
)

var dataElementsCmd = &cobra.Command{
	Use:   "data-elements",
	Short: "Manage Data Elements",
	Long:  `Commands for managing Data Elements resources.`,
}

func init() {
	datasourcesCmd.AddCommand(dataElementsCmd)
}
