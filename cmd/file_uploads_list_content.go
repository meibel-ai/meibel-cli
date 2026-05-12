package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	fileUploadsListContentPrefix string
	fileUploadsListContentContinuationToken string
	fileUploadsListContentLimit int64
)

var fileUploadsListContentCmd = &cobra.Command{
	Use:   "list-content <datasource-id>",
	Short: "List Content",
	Long:  `List Content

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources file-uploads list-content <datasource-id> --prefix=<value> --continuation-token=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		opts := &sdk.FileUploadsListContentOptions{}
		if fileUploadsListContentPrefix != "" {
			opts.Prefix = &fileUploadsListContentPrefix
		}
		if fileUploadsListContentContinuationToken != "" {
			opts.ContinuationToken = &fileUploadsListContentContinuationToken
		}
		if fileUploadsListContentLimit != 0 {
			opts.Limit = &fileUploadsListContentLimit
		}

		iter := client.Datasources.FileUploads.ListContent(ctx, datasourceId, opts)

		var items []interface{}
		for iter.Next(ctx) {
			items = append(items, iter.Item())
		}
		if err := iter.Err(); err != nil {
			return err
		}

		return output.Print(items)
	},
}

func init() {
	fileUploadsCmd.AddCommand(fileUploadsListContentCmd)

	fileUploadsListContentCmd.Flags().StringVarP(&fileUploadsListContentPrefix, "prefix", "", "", "Filter content by path prefix")
	fileUploadsListContentCmd.Flags().StringVarP(&fileUploadsListContentContinuationToken, "continuation-token", "", "", "Token for pagination")
	fileUploadsListContentCmd.Flags().Int64VarP(&fileUploadsListContentLimit, "limit", "", 1000, "Maximum items to return")
}
