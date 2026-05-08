package cmd

import (
	"github.com/spf13/cobra"
)

var fileUploadsCmd = &cobra.Command{
	Use:   "file-uploads",
	Short: "Manage FileUploads",
	Long:  `Commands for managing FileUploads resources.`,
}

func init() {
	datasourcesCmd.AddCommand(fileUploadsCmd)
}
