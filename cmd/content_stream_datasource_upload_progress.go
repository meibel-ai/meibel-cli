package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var contentStreamDatasourceUploadProgressCmd = &cobra.Command{
	Use:   "stream-datasource-upload-progress <datasource-id> <upload-id>",
	Short: "Stream upload progress events (legacy)",
	Long:  `Subscribe to real-time upload progress updates via Server-Sent Events. Consider using /uploads/{upload_id}/progress instead.

Arguments:
  datasource-id: required
  upload-id: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel datasources content stream-datasource-upload-progress <datasource-id> <upload-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		uploadId := args[1]

		err := client.Datasources.Content.StreamDatasourceUploadProgress(ctx, datasourceId, uploadId)
		if err != nil {
			return err
		}

		fmt.Println("Success")
		return nil
	},
}

func init() {
	contentCmd.AddCommand(contentStreamDatasourceUploadProgressCmd)

}
