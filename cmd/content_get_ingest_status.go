package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel/internal/output"
)

var contentGetIngestStatusCmd = &cobra.Command{
	Use:   "get-ingest-status <datasource-id>",
	Short: "Get Ingest Status",
	Long:  `Get Ingest Status

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources content get-ingest-status <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		result, err := client.Datasources.Content.GetIngestStatus(ctx, datasourceId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	contentCmd.AddCommand(contentGetIngestStatusCmd)

}
