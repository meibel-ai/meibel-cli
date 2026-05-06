package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var contentGetDatasourceContentMetadataCmd = &cobra.Command{
	Use:   "get-datasource-metadata <datasource-id> <path>",
	Short: "Get content metadata",
	Long:  `Get metadata information for a file or directory in the datasource

Arguments:
  datasource-id: required
  path: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel datasources content get-datasource-metadata <datasource-id> <path>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		path := args[1]

		result, err := client.Datasources.Content.GetDatasourceContentMetadata(ctx, datasourceId, path)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	contentCmd.AddCommand(contentGetDatasourceContentMetadataCmd)

}
