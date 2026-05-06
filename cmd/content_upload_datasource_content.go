package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var contentUploadDatasourceContentCmd = &cobra.Command{
	Use:   "upload-datasource <datasource-id>",
	Short: "Upload Content",
	Long:  `Proxy upload with zero-copy streaming.

This endpoint maintains the multipart form data structure and streams
it directly to the backend service without buffering files in memory.
The multipart parsing happens on the backend service side.

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources content upload-datasource <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		result, err := client.Datasources.Content.UploadDatasourceContent(ctx, datasourceId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	contentCmd.AddCommand(contentUploadDatasourceContentCmd)

}
