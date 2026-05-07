package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	fileUploadListContentPrefix string
	fileUploadListContentContinuationToken string
	fileUploadListContentLimit int64
)

var fileUploadListContentCmd = &cobra.Command{
	Use:   "list-content <datasource-id>",
	Short: "List Content",
	Long:  `List Content

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources file-upload list-content <datasource-id> --prefix=<value> --continuation-token=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		opts := &sdk.ListContentOptions{}
		if fileUploadListContentPrefix != "" {
			opts.Prefix = &fileUploadListContentPrefix
		}
		if fileUploadListContentContinuationToken != "" {
			opts.ContinuationToken = &fileUploadListContentContinuationToken
		}
		if fileUploadListContentLimit != 0 {
			opts.Limit = &fileUploadListContentLimit
		}

		iter := client.Datasources.FileUpload.ListContent(ctx, datasourceId, opts)

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
	fileUploadCmd.AddCommand(fileUploadListContentCmd)

	fileUploadListContentCmd.Flags().StringVarP(&fileUploadListContentPrefix, "prefix", "", "", "Filter content by path prefix")
	fileUploadListContentCmd.Flags().StringVarP(&fileUploadListContentContinuationToken, "continuation-token", "", "", "Token for pagination")
	fileUploadListContentCmd.Flags().Int64VarP(&fileUploadListContentLimit, "limit", "", 1000, "Maximum items to return")
}
