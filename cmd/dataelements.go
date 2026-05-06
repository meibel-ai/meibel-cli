package cmd

import (
	"github.com/spf13/cobra"
)

var dataelementsCmd = &cobra.Command{
	Use:   "dataelements",
	Short: "Manage dataelements",
	Long:  `Commands for managing dataelements resources.`,
}

func init() {
	datasourcesCmd.AddCommand(dataelementsCmd)
}
