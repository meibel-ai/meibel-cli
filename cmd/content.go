package cmd

import (
	"github.com/spf13/cobra"
)

var contentCmd = &cobra.Command{
	Use:   "content",
	Short: "Manage content",
	Long:  `Commands for managing content resources.`,
}

func init() {
	datasourcesCmd.AddCommand(contentCmd)
}
