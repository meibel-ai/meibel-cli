package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var contentDownloadDatasourceContentCmd = &cobra.Command{
	Use:   "download-datasource <datasource-id> <path>",
	Short: "Download content file",
	Long:  `Download a file from the datasource with streaming support for large files

Arguments:
  datasource-id: required
  path: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel datasources content download-datasource <datasource-id> <path>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		path := args[1]

		err := client.Datasources.Content.DownloadDatasourceContent(ctx, datasourceId, path)
		if err != nil {
			return err
		}

		fmt.Println("Success")
		return nil
	},
}

func init() {
	contentCmd.AddCommand(contentDownloadDatasourceContentCmd)

}
