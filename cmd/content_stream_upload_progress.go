package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var contentStreamUploadProgressCmd = &cobra.Command{
	Use:   "stream-upload-progress <upload-id>",
	Short: "Stream Upload Progress",
	Long:  `Stream Upload Progress

Arguments:
  upload-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources content stream-upload-progress <upload-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		uploadId := args[0]

		err := client.Datasources.Content.StreamUploadProgress(ctx, uploadId)
		if err != nil {
			return err
		}

		fmt.Println("Success")
		return nil
	},
}

func init() {
	contentCmd.AddCommand(contentStreamUploadProgressCmd)

}
