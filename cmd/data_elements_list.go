package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	dataElementsListCursor string
	dataElementsListLimit int64
)

var dataElementsListCmd = &cobra.Command{
	Use:   "list <datasource-id>",
	Short: "List Data Elements",
	Long:  `List Data Elements

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources data-elements list <datasource-id> --cursor=<value> --limit=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		opts := &sdk.DataElementsListOptions{}
		if dataElementsListCursor != "" {
			opts.Cursor = &dataElementsListCursor
		}
		if dataElementsListLimit != 0 {
			opts.Limit = &dataElementsListLimit
		}

		iter := client.Datasources.DataElements.List(ctx, datasourceId, opts)

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
	dataElementsCmd.AddCommand(dataElementsListCmd)

	dataElementsListCmd.Flags().StringVarP(&dataElementsListCursor, "cursor", "", "", "Cursor for pagination")
	dataElementsListCmd.Flags().Int64VarP(&dataElementsListLimit, "limit", "", 100, "Maximum items to return")
}
