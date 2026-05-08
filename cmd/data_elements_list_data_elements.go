package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	dataElementsListDataElementsCursor string
	dataElementsListDataElementsLimit int64
)

var dataElementsListDataElementsCmd = &cobra.Command{
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

		opts := &sdk.ListDataElementsOptions{}
		if dataElementsListDataElementsCursor != "" {
			opts.Cursor = &dataElementsListDataElementsCursor
		}
		if dataElementsListDataElementsLimit != 0 {
			opts.Limit = &dataElementsListDataElementsLimit
		}

		iter := client.Datasources.DataElements.ListDataElements(ctx, datasourceId, opts)

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
	dataElementsCmd.AddCommand(dataElementsListDataElementsCmd)

	dataElementsListDataElementsCmd.Flags().StringVarP(&dataElementsListDataElementsCursor, "cursor", "", "", "Cursor for pagination")
	dataElementsListDataElementsCmd.Flags().Int64VarP(&dataElementsListDataElementsLimit, "limit", "", 100, "Maximum items to return")
}
