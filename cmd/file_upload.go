package cmd

import (
	"github.com/spf13/cobra"
)

var fileUploadCmd = &cobra.Command{
	Use:   "file-upload",
	Short: "Manage File Upload",
	Long:  `Commands for managing File Upload resources.`,
}

func init() {
	datasourcesCmd.AddCommand(fileUploadCmd)
}
