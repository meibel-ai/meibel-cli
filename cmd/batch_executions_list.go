package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	batchExecutionsListInputDatasourceId string
	batchExecutionsListOffset int64
	batchExecutionsListLimit string
	batchExecutionsListSortBy string
	batchExecutionsListSortOrder string
)

var batchExecutionsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Batch Executions",
	Long:  `List Batch Executions`,
	Example: "meibel batch-executions list --input-datasource-id=<value> --offset=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		opts := &sdk.BatchExecutionsListOptions{}
		if batchExecutionsListInputDatasourceId != "" {
			opts.InputDatasourceId = &batchExecutionsListInputDatasourceId
		}
		if batchExecutionsListOffset != 0 {
			opts.Offset = &batchExecutionsListOffset
		}
		if batchExecutionsListLimit != "" {
			opts.Limit = &batchExecutionsListLimit
		}
		if batchExecutionsListSortBy != "" {
			opts.SortBy = &batchExecutionsListSortBy
		}
		if batchExecutionsListSortOrder != "" {
			opts.SortOrder = &batchExecutionsListSortOrder
		}

		iter := client.BatchExecutions.List(ctx, opts)

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
	batchExecutionsCmd.AddCommand(batchExecutionsListCmd)

	batchExecutionsListCmd.Flags().StringVarP(&batchExecutionsListInputDatasourceId, "input-datasource-id", "", "", "Filter by input datasource ID")
	batchExecutionsListCmd.Flags().Int64VarP(&batchExecutionsListOffset, "offset", "", 0, "The offset parameter")
	batchExecutionsListCmd.Flags().StringVarP(&batchExecutionsListLimit, "limit", "", "", "The limit parameter")
	batchExecutionsListCmd.Flags().StringVarP(&batchExecutionsListSortBy, "sort-by", "", "start_time", "Field to sort by: start_time, status")
	batchExecutionsListCmd.Flags().StringVarP(&batchExecutionsListSortOrder, "sort-order", "", "desc", "The sort-order parameter")
}
