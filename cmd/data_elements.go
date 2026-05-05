package cmd

import (
	"github.com/spf13/cobra"
)

var dataElementsCmd = &cobra.Command{
	Use:   "data-elements",
	Short: "Manage DataElements",
	Long:  `Commands for managing DataElements resources.`,
}

func init() {
	rootCmd.AddCommand(dataElementsCmd)
}
