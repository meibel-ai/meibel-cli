package cmd

import (
	"context"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/tui"
)

var fileUploadStreamUploadProgressCmd = &cobra.Command{
	Use:   "stream-progress <upload-id>",
	Short: "Stream Upload Progress",
	Long:  `Stream Upload Progress

Arguments:
  upload-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources file-upload stream-progress <upload-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		uploadId := args[0]

		// Set up signal handling for graceful shutdown
		ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
		defer cancel()

		stream, err := client.Datasources.FileUpload.StreamUploadProgress(ctx, uploadId)
		if err != nil {
			return err
		}
		defer stream.Close()

		return tui.StreamEvents(ctx, stream)
	},
}

func init() {
	fileUploadCmd.AddCommand(fileUploadStreamUploadProgressCmd)

}
