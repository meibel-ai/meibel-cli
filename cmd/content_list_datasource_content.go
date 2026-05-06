package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	contentListDatasourceContentPrefix string
	contentListDatasourceContentContinuationToken string
	contentListDatasourceContentLimit int64
)

var contentListDatasourceContentCmd = &cobra.Command{
	Use:   "list-datasource <datasource-id>",
	Short: "List datasource content",
	Long:  `List files and directories in a datasource with optional filtering and pagination

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources content list-datasource <datasource-id> --prefix=<value> --continuation-token=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		opts := &sdk.ListDatasourceContentOptions{}
		if contentListDatasourceContentPrefix != "" {
			opts.Prefix = &contentListDatasourceContentPrefix
		}
		if contentListDatasourceContentContinuationToken != "" {
			opts.ContinuationToken = &contentListDatasourceContentContinuationToken
		}
		if contentListDatasourceContentLimit != 0 {
			opts.Limit = &contentListDatasourceContentLimit
		}

		result, err := client.Datasources.Content.ListDatasourceContent(ctx, datasourceId, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	contentCmd.AddCommand(contentListDatasourceContentCmd)

	contentListDatasourceContentCmd.Flags().StringVarP(&contentListDatasourceContentPrefix, "prefix", "", "", "Filter content by path prefix")
	contentListDatasourceContentCmd.Flags().StringVarP(&contentListDatasourceContentContinuationToken, "continuation-token", "", "", "Token for pagination to get next page of results")
	contentListDatasourceContentCmd.Flags().Int64VarP(&contentListDatasourceContentLimit, "limit", "", 1000, "Maximum number of items to return (1-10000)")
}
