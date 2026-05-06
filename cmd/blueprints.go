package cmd

import (
	"github.com/spf13/cobra"
)

var blueprintsCmd = &cobra.Command{
	Use:   "blueprints",
	Short: "Manage blueprints",
	Long:  `Commands for managing blueprints resources.`,
}

func init() {
	rootCmd.AddCommand(blueprintsCmd)
}
