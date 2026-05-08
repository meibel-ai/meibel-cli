package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var fileUploadsGetIngestStatusCmd = &cobra.Command{
	Use:   "get-ingest-status <datasource-id>",
	Short: "Get Ingest Status",
	Long:  `Get Ingest Status

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources file-uploads get-ingest-status <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		result, err := client.Datasources.FileUploads.GetIngestStatus(ctx, datasourceId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	fileUploadsCmd.AddCommand(fileUploadsGetIngestStatusCmd)

}
