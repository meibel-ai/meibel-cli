package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	contentListContentPrefix string
	contentListContentContinuationToken string
	contentListContentLimit int64
)

var contentListContentCmd = &cobra.Command{
	Use:   "list <datasource-id>",
	Short: "List Content",
	Long:  `List Content

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources content list <datasource-id> --prefix=<value> --continuation-token=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		opts := &sdk.ListContentOptions{}
		if contentListContentPrefix != "" {
			opts.Prefix = &contentListContentPrefix
		}
		if contentListContentContinuationToken != "" {
			opts.ContinuationToken = &contentListContentContinuationToken
		}
		if contentListContentLimit != 0 {
			opts.Limit = &contentListContentLimit
		}

		result, err := client.Datasources.Content.ListContent(ctx, datasourceId, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	contentCmd.AddCommand(contentListContentCmd)

	contentListContentCmd.Flags().StringVarP(&contentListContentPrefix, "prefix", "", "", "Filter content by path prefix")
	contentListContentCmd.Flags().StringVarP(&contentListContentContinuationToken, "continuation-token", "", "", "Token for pagination")
	contentListContentCmd.Flags().Int64VarP(&contentListContentLimit, "limit", "", 1000, "Maximum items to return")
}
